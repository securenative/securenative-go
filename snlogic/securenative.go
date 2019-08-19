package snlogic

import (
	"errors"
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

const signatureHeader = "x-securenative"

func NewSecureNative(apiKey string, options *models.SecureNativeOptions) (SDK, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("You must pass your SecureNative api key")
	}

	if options == nil {
		options = &models.SecureNativeOptions{}
	}
	manger := NewSnEventManger(apiKey, options)
	return &Securenative{ApiKey: apiKey, SnOptions: options, EventManager: manger}, nil
}

func (this *Securenative) Verify(e models.SnEvent) *models.RiskResult {
	return this.EventManager.SendSync(e, this.SnOptions.ApiUrl+"/verify")
}

func (this *Securenative) Track(e models.SnEvent) {
	this.EventManager.SendAsync(e, this.SnOptions.ApiUrl+"/track")
}

func (this *Securenative) Flow(flowId int64, e models.SnEvent) *models.RiskResult {
	return this.EventManager.SendSync(e, this.SnOptions.ApiUrl+"/flow")
}

func (this *Securenative) IsRequestFromSn(snHeaderValue string, payload string) bool {
	if len(snHeaderValue) == 0 || len(payload) == 0 {
		return false
	}
	encrypted, err := Encrypt(string(payload), this.ApiKey)
	if err == nil {
		if encrypted == snHeaderValue {
			return true
		}
	}
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
