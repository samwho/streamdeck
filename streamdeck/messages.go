package streamdeck

import (
	"context"
	"encoding/json"
)

type LogMessagePayload struct {
	Message string `json:"message"`
}

func NewLogMessage(message string) Event {
	return NewEvent(nil, LogMessage, LogMessagePayload{Message: message})
}

type RegisterEvent struct {
	Event string `json:"event"`
	UUID  string `json:"uuid"`
}

func NewRegisterEvent(params RegistrationParams) RegisterEvent {
	return RegisterEvent{
		Event: params.RegisterEvent,
		UUID:  params.PluginUUID,
	}
}

type Event struct {
	Action  string `json:"action,omitempty"`
	Event   string `json:"event,omitempty"`
	Context string `json:"context,omitempty"`
	Device  string `json:"device,omitempty"`
	Payload string `json:"payload,omitempty"`
}

func NewEvent(ctx context.Context, name string, payload interface{}) Event {
	payloadStr, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return Event{
		Event:   name,
		Context: getContext(ctx),
		Payload: string(payloadStr),
	}
}
