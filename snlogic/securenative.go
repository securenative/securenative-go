package snlogic

import (
	"errors"
	"fmt"
	"github.com/securenative/securenative-go/models"
)

type SDK interface {
	Verify(e models.SnEvent) *models.RiskResult
	Track(e models.SnEvent)
	Flow(flowId int64, e models.SnEvent) *models.RiskResult
	IsRequestFromSn(snHeaderValue string, payload string) bool
}

type Securenative struct {
	EventManager EventManager
	SnOptions    *models.SecureNativeOptions
	ApiKey       string
}

var instance *Securenative

const signatureHeader = "x-securenative"

func Init(apiKey string, options *models.SecureNativeOptions) (SDK, error) {
	if instance != nil {
		return nil, errors.New("This SDK was already initialized")
	}

	if len(apiKey) == 0 {
		return nil, errors.New("You must pass your SecureNative api key")
	}

	if options == nil {
		options = &models.SecureNativeOptions{}
	}
	manger := NewSnEventManger(apiKey, options)
	IsLogEnabledFlag = options.IsLoggingEnabled
	instance = &Securenative{ApiKey: apiKey, SnOptions: options, EventManager: manger}
	SnLog("SN SDK was initialized")
	return instance, nil
}

func GetInstance() (*Securenative, error) {
	if instance == nil {
		SnLog("SN SDK wasn't initialized")
		return nil, errors.New("You should call init(api_key, options) before making any other sdk function call")
	}
	return instance, nil
}

func (this *Securenative) Verify(e models.SnEvent) *models.RiskResult {
	SnLog(fmt.Sprintf("verify event call %v ", e))
	return this.EventManager.SendSync(e, this.SnOptions.ApiUrl+"/verify")
}

func (this *Securenative) Track(e models.SnEvent) {
	SnLog(fmt.Sprintf("track event call %v ", e))
	this.EventManager.SendAsync(e, this.SnOptions.ApiUrl+"/track")
}

func (this *Securenative) Flow(flowId int64, e models.SnEvent) *models.RiskResult {
	SnLog(fmt.Sprintf("flow event call %v ", e))
	return this.EventManager.SendSync(e, this.SnOptions.ApiUrl+"/flow")
}

func (this *Securenative) IsRequestFromSn(snHeaderValue string, payload string) bool {
	SnLog("verifying request from SN ")

	if len(snHeaderValue) == 0 || len(payload) == 0 {
		SnLog("payload or header are empty, verification failed")
		return false
	}
	encrypted, err := Encrypt(string(payload), this.ApiKey)
	if err == nil {
		if encrypted == snHeaderValue {
			SnLog("request was verified")
			return true
		}
	}
	SnLog("verification failed")
	return false
}

/*func (this *Securenative) SnMiddleware(ctx *gin.Context) {
	if len(ctx.Request.Header[signatureHeader]) > 0 && len(ctx.Request.Header[signatureHeader][0]) > 0{
		signature := ctx.Request.Header[signatureHeader][0]
		if len(signature) > 0 {
			payload, err := ioutil.ReadAll(ctx.Request.Body)
			if err == nil {
				encrypted, err := Encrypt(string(payload), this.ApiKey)
				if err == nil {
					if encrypted == signature {
						ctx.Next()
						return
					}

				}
			}
		}
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
}
*/
