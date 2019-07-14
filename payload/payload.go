package payload

import (
	"encoding/json"
)

type Target int

const (
	HardwareAndSoftware Target = 0
	OnlyHardware        Target = 1
	OnlySoftware        Target = 2
)

type LogMessage struct {
	Message string `json:"message"`
}

type OpenURL struct {
	URL string `json:"url"`
}

type SetTitle struct {
	Title  string `json:"title"`
	Target Target `json:"target"`
}

type SetImage struct {
	Base64Image string `json:"image"`
	Target      Target `json:"target"`
}

type SetState struct {
	State int `json:"state"`
}

type SwitchProfile struct {
	Profile string `json:"profile"`
}

type DidReceiveSettings struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	IsInMultiAction bool            `json:"isInMultiAction,omitempty"`
}

type Coordinates struct {
	Column int `json:"column,omitempty"`
	Row    int `json:"row,omitempty"`
}

type DidReceiveGlobalSettings struct {
	Settings json.RawMessage `json:"settings,omitempty"`
}

type KeyDown struct {
	Settings         json.RawMessage `json:"settings,omitempty"`
	Coordinates      Coordinates     `json:"coordinates,omitempty"`
	State            int             `json:"state,omitempty"`
	UserDesiredState int             `json:"userDesiredState,omitempty"`
	IsInMultiAction  bool            `json:"isInMultiAction,omitempty"`
}

type KeyUp struct {
	Settings         json.RawMessage `json:"settings,omitempty"`
	Coordinates      Coordinates     `json:"coordinates,omitempty"`
	State            int             `json:"state,omitempty"`
	UserDesiredState int             `json:"userDesiredState,omitempty"`
	IsInMultiAction  bool            `json:"isInMultiAction,omitempty"`
}

type WillAppear struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	State           int             `json:"state,omitempty"`
	IsInMultiAction bool            `json:"isInMultiAction,omitempty"`
}

type WillDisappear struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	State           int             `json:"state,omitempty"`
	IsInMultiAction bool            `json:"isInMultiAction,omitempty"`
}

type TitleParametersDidChange struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	State           int             `json:"state,omitempty"`
	Title           string          `json:"title,omitempty"`
	TitleParameters TitleParameters `json:"titleParameters,omitempty"`
}

type TitleParameters struct {
	FontFamily     string `json:"fontFamily,omitempty"`
	FontSize       int    `json:"fontSize,omitempty"`
	FontStyle      string `json:"fontStyle,omitempty"`
	FontUnderline  bool   `json:"fontUnderline,omitempty"`
	ShowTitle      bool   `json:"showTitle,omitempty"`
	TitleAlignment string `json:"titleAlignment,omitempty"`
	TitleColor     string `json:"titleColor,omitempty"`
}

type ApplicationDidLaunch struct {
	Application string `json:"application,omitempty"`
}

type ApplicationDidTerminate struct {
	Application string `json:"application,omitempty"`
}
