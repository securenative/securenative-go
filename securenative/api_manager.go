package securenative

import (
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
	// TODO: Debug LOG - "Track event call"
	event := NewSDKEvent(eventOptions, m.Options)
	m.EventManager.SendAsync(event, ApiRoute.Track)
}

func (m *ApiManager) Verify(eventOptions EventOptions) VerifyResult {
	// TODO: Debug LOG - "Verify event call"
	event := NewSDKEvent(eventOptions, m.Options)

	res, err := m.EventManager.SendSync(event, ApiRoute.Verify, false)
	if err != nil {
		// TODO: Debug LOG - "Failed to call verify; {}"
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
