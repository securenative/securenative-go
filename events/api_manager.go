package events

import (
	"fmt"
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/enums"
	"github.com/securenative/securenative-go/models"
	"github.com/securenative/securenative-go/utils"
)

type ApiManagerInterface interface {
	Track(eventOptions models.EventOptions) error
	Verify(eventOptions models.EventOptions) (models.VerifyResult, error)
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

func (m *ApiManager) Track(eventOptions models.EventOptions) error {
	logger := utils.GetLogger()
	logger.Debug("Track event call")

	event, err := models.NewSDKEvent(eventOptions, m.Options)
	if err != nil {
		return err
	}

	m.EventManager.SendAsync(event, enums.ApiRoute.Track)
	return nil
}

func (m *ApiManager) Verify(eventOptions models.EventOptions) (*models.VerifyResult, error) {
	logger := utils.GetLogger()
	logger.Debug("Verify event call")

	event, err := models.NewSDKEvent(eventOptions, m.Options)

	if err != nil {
		return nil, err
	}

	res, err := m.EventManager.SendSync(event, enums.ApiRoute.Verify, false)
	if err != nil || res == nil {
		logger.Debug(fmt.Sprintf("Failed to call verify; %s", err))
		if m.Options.FailOverStrategy == enums.FailOverStrategy.FailOpen {
			return &models.VerifyResult{RiskLevel: enums.RiskLevel.Low, Score: 0, Triggers: nil}, nil
		}
		return &models.VerifyResult{RiskLevel: enums.RiskLevel.High, Score: 1, Triggers: nil}, nil
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
	}, nil
}
