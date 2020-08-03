package models

type EventInput struct {
	RequestID  string                 `json:"rid"`
	EventType  string                 `json:"eventType"`
	UserID     interface{}            `json:"userId"`
	UserTraits UserTraits             `json:"userTraits"`
	Request    RequestContext         `json:"request"`
	Properties map[string]interface{} `json:"properties"`
	Timestamp  string                 `json:"timestamp"`
}
