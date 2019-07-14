package streamdeck

import (
	"context"
	"encoding/json"

	sdcontext "github.com/samwho/streamdeck/context"
)

type Event struct {
	Action     string          `json:"action,omitempty"`
	Event      string          `json:"event,omitempty"`
	UUID       string          `json:"uuid,omitempty"`
	Context    string          `json:"context,omitempty"`
	Device     string          `json:"device,omitempty"`
	DeviceInfo DeviceInfo      `json:"deviceInfo,omitempty"`
	Payload    json.RawMessage `json:"payload,omitempty"`
}

type DeviceInfo struct {
	DeviceName string     `json:"deviceName,omitempty"`
	Type       DeviceType `json:"type,omitempty"`
	Size       DeviceSize `json:"size,omitempty"`
}

type DeviceSize struct {
	Columns int `json:"columns,omitempty"`
	Rows    int `json:"rows,omitempty"`
}

type DeviceType int

const (
	StreamDeck       DeviceType = 0
	StreamDeckMini   DeviceType = 1
	StreamDeckXL     DeviceType = 2
	StreamDeckMobile DeviceType = 3
)

func NewEvent(ctx context.Context, name string, payload interface{}) Event {
	p, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return Event{
		Event:   name,
		Action:  sdcontext.Action(ctx),
		Context: sdcontext.Context(ctx),
		Device:  sdcontext.Device(ctx),
		Payload: p,
	}
}
