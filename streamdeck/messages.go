package streamdeck

import (
	"context"
)

type Event struct {
	Action  string      `json:"action,omitempty"`
	Event   string      `json:"event,omitempty"`
	UUID    string      `json:"uuid,omitempty"`
	Context string      `json:"context,omitempty"`
	Device  string      `json:"device,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewEvent(ctx context.Context, name string, payload interface{}) Event {
	return Event{
		Event:   name,
		Context: getContext(ctx),
		Payload: payload,
	}
}

type LogMessagePayload struct {
	Message string `json:"message"`
}

type OpenURLPayload struct {
	URL string `json:"url"`
}

type Target int

const (
	HardwareAndSoftware Target = 0
	OnlyHardware        Target = 1
	OnlySoftware        Target = 2
)

type SetTitlePayload struct {
	Title  string `json:"title"`
	Target Target `json:"target"`
}

type SetImagePayload struct {
	Base64Image string `json:"image"`
	Target      Target `json:"target"`
}
