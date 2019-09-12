package snlogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/securenative/securenative-go/models"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const authorization = "Authorization"
const useragent = "User-Agent"
const useragentValue = "securenative-go.snlogic.SecureNative-go"
const snVersion = "SN-Version"
const acceptHeader = "Accept"
const acceptValue = "application/json"

var defaultRiskResult = models.RiskResult{"low", 0.0, []string{}}

type EventManager interface {
	SendSync(e models.SnEvent, requestUrl string) *models.RiskResult
	SendAsync(e models.SnEvent, requestUrl string)
	SendEventsFromChannel()
}

type SnEventManger struct {
	userAgentValue string `envDefault:"com.securenative.snlogic.SecureNative-java"`
	snVersion      string `envDefault:"SN-version"`
	ApiKey         string `envDefault:""`
	client         *http.Client
	req            *http.Request
	events         chan models.Message
	options        *models.SecureNativeOptions
}

func NewSnEventManger(apiKey string, options *models.SecureNativeOptions) (snEventManager EventManager) {
	SnLog(fmt.Sprintf("intitalizing event manager"))

	tempClient := http.Client{
		Timeout: time.Duration(options.Timeout),
	}
	tempEvents := make(chan models.Message)
	snEventManager = &SnEventManger{ApiKey: apiKey, client: &tempClient, events: tempEvents, options: options}

	if options != nil && options.IsSdkEnabled && options.AutoSend {
		SnLog(fmt.Sprintf("Starting event auto send mechanism"))
		go snEventManager.SendEventsFromChannel()
	}
	return
}

func (this *SnEventManger) SendEventsFromChannel() {

	for {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		m := <-this.events
		SnLog(fmt.Sprintf("Sending async event %v ", m))
		this.SendAsync(m.Event, m.Url)
	}
}

func (this *SnEventManger) SendSync(e models.SnEvent, requestUrl string) (riskResult *models.RiskResult) {
	if this.options != nil && !this.options.IsSdkEnabled {
		return &defaultRiskResult
	}
	req, err := createRequest(e, requestUrl, this.ApiKey)
	if err != nil {
		return &defaultRiskResult
	}
	SnLog(fmt.Sprintf("Sending sync event %v ", e))
	resp, err := this.client.Do(req)
	if err != nil {
		return &defaultRiskResult
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &defaultRiskResult
	}
	err = json.Unmarshal(body, riskResult)
	if err != nil {
		return &defaultRiskResult
	}
	SnLog(fmt.Sprintf("Response to sync event %v i %v", e, riskResult))
	return
}

func (this *SnEventManger) SendAsync(e models.SnEvent, requestUrl string) {
	if this.options != nil && !this.options.IsSdkEnabled {
		return
	}
	req, err := createRequest(e, requestUrl, this.ApiKey)
	if err != nil && this.options != nil && !this.options.AutoSend {
		this.events <- models.Message{e, requestUrl}
		return
	}
	_, err = this.client.Do(req)
	if err != nil && this.options != nil && !this.options.AutoSend {
		this.events <- models.Message{e, requestUrl}
		return
	}
}

func getVersion() string {
	return "version" //TODO: Add dynamic version from
}

func createRequest(e models.SnEvent, requestUrl string, apiKey string) (req *http.Request, err error) {
	payload, _ := json.Marshal(e)
	req, err = http.NewRequest("POST", requestUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set(authorization, apiKey)
	req.Header.Set(useragent, useragentValue)
	req.Header.Set(acceptHeader, acceptValue)
	req.Header.Set(snVersion, getVersion())
	return req, nil

}
