package events

import (
	"fmt"
	"github.com/securenative/securenative-go/securenative/config"
	"github.com/securenative/securenative-go/securenative/enums"
	"github.com/securenative/securenative-go/securenative/models"
	"github.com/securenative/securenative-go/securenative/utils"
)

type ApiManagerInterface interface {
	Track(eventOptions models.EventOptions)
	Verify(eventOptions models.EventOptions) models.VerifyResult
}

type ApiManager struct {
	EventManager *EventManager
	Options      config.SecureNativeOptions
}

func NewApiManager(eventManager *EventManager, options config.SecureNativeOptions) *ApiManager {
	return &ApiManager{
		EventManager: eventManager,
		Options:      options,
	}
}

func (m *ApiManager) Track(eventOptions models.EventOptions) {
	logger := utils.GetLogger()
	logger.Debug("Track event call")

	event := models.NewSDKEvent(eventOptions, m.Options)
	m.EventManager.SendAsync(event, enums.ApiRoute.Track)
}

func (m *ApiManager) Verify(eventOptions models.EventOptions) *models.VerifyResult {
	logger := utils.GetLogger()
	logger.Debug("Verify event call")

	event := models.NewSDKEvent(eventOptions, m.Options)

	res, err := m.EventManager.SendSync(event, enums.ApiRoute.Verify, false)
	if err != nil {
		logger.Debug(fmt.Sprintf("Failed to call verify; %s", err))
		if m.Options.FailOverStrategy == enums.FailOverStrategy.FailOpen {
			return &models.VerifyResult{RiskLevel: enums.RiskLevel.Low, Score: 0, Triggers: nil}
		}
		return &models.VerifyResult{RiskLevel: enums.RiskLevel.High, Score: 1, Triggers: nil}
	}

	score := res["score"].(float64)

	var triggers []string
	for _, v := range res["triggers"].([]interface{}) {
		triggers = append(triggers, fmt.Sprint(v))
	}

	return &models.VerifyResult{
		RiskLevel: res["riskLevel"].(string),
		Score:     score,
		Triggers:  triggers,
	}
}
