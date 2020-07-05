package securenative

import (
	"fmt"
	. "github.com/securenative/securenative-go/securenative/config"
	. "github.com/securenative/securenative-go/securenative/enums"
	. "github.com/securenative/securenative-go/securenative/models"
	"strconv"
	"strings"
)

type ApiManagerInterface interface {
	Track(eventOptions EventOptions)
	Verify(eventOptions EventOptions) VerifyResult
}

type ApiManager struct {
	EventManager *EventManager
	Options      SecureNativeOptions
}

func NewApiManager(eventManager *EventManager, options SecureNativeOptions) *ApiManager {
	return &ApiManager{
		EventManager: eventManager,
		Options:      options,
	}
}

func (m *ApiManager) Track(eventOptions EventOptions) {
	logger := GetLogger()
	logger.Debug("Track event call")

	event := NewSDKEvent(eventOptions, m.Options)
	m.EventManager.SendAsync(event, ApiRoute.Track)
}

func (m *ApiManager) Verify(eventOptions EventOptions) VerifyResult {
	logger := GetLogger()
	logger.Debug("Verify event call")

	event := NewSDKEvent(eventOptions, m.Options)

	res, err := m.EventManager.SendSync(event, ApiRoute.Verify, false)
	if err != nil {
		logger.Debug(fmt.Sprintf("Failed to call verify; %s", err))
		if m.Options.FailOverStrategy == FailOverStrategy.FailOpen {
			return VerifyResult{RiskLevel: RiskLevel.Low, Score: 0, Triggers: nil}
		}
		return VerifyResult{RiskLevel: RiskLevel.High, Score: 1, Triggers: nil}
	}

	score, _ := strconv.Atoi(res["score"])
	triggers :=  strings.Split(res["triggers"], ",")
	return VerifyResult{
		RiskLevel: res["riskLevel"],
		Score:     int32(score),
		Triggers:  triggers,
	}
}
