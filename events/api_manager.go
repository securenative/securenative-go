package events

import (
	"fmt"
	"github.com/securenative/securenative-go"
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/enums"
	"github.com/securenative/securenative-go/models"
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
	logger := securenative_go.GetLogger()
	logger.Debug("Track event call")

	event, err := models.NewSDKEvent(eventOptions, m.Options)
	if err != nil {
		return err
	}

	m.EventManager.SendAsync(event, enums.ApiRoute.Track)
	return nil
}

func (m *ApiManager) Verify(eventOptions models.EventOptions) (*models.VerifyResult, error) {
	logger := securenative_go.GetLogger()
	logger.Debug("Verify event call")

	event, err := models.NewSDKEvent(eventOptions, m.Options)

	if err != nil {
		return nil, err
	}

	res, err := m.EventManager.SendSync(event, enums.ApiRoute.Verify)
	if err != nil || res == nil {
		logger.Debug(fmt.Sprintf("Failed to call verify; %s", err))
		if m.Options.FailOverStrategy == enums.FailOverStrategy.FailOpen {
			return &models.VerifyResult{RiskLevel: enums.RiskLevel.Low, Score: 0, Triggers: []string{}}, nil
		}
		return &models.VerifyResult{RiskLevel: enums.RiskLevel.High, Score: 1, Triggers: []string{}}, nil
	}

	var riskLevel string
	if res["riskLevel"] != nil {
		riskLevel = res["riskLevel"].(string)
	} else {
		if m.Options.FailOverStrategy == enums.FailOverStrategy.FailOpen {
			riskLevel = enums.RiskLevel.Low
		} else {
			riskLevel = enums.RiskLevel.High
		}
	}

	var score float64
	if res["score"] != nil {
		score = res["score"].(float64)
	} else {
		if m.Options.FailOverStrategy == enums.FailOverStrategy.FailOpen {
			score = 0
		} else {
			score = 1
		}
	}

	var triggers []string
	if res["triggers"] != nil {
		for _, v := range res["triggers"].([]interface{}) {
			triggers = append(triggers, fmt.Sprint(v))
		}
	}

	return &models.VerifyResult{
		RiskLevel: riskLevel,
		Score:     score,
		Triggers:  triggers,
	}, nil
}
