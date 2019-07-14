package streamdeck

import "encoding/json"

type LogMessagePayload struct {
	Message string `json:"message"`
}

type OpenURLPayload struct {
	URL string `json:"url"`
}

type SetTitlePayload struct {
	Title  string `json:"title"`
	Target Target `json:"target"`
}

type SetImagePayload struct {
	Base64Image string `json:"image"`
	Target      Target `json:"target"`
}

type SetStatePayload struct {
	State int `json:"state"`
}

type SwitchProfilePayload struct {
	Profile string `json:"profile"`
}

type DidReceiveSettingsPayload struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	IsInMultiAction bool            `json:"isInMultiAction,omitempty"`
}

type Coordinates struct {
	Column int `json:"column,omitempty"`
	Row    int `json:"row,omitempty"`
}

type DidReceiveGlobalSettingsPayload struct {
	Settings json.RawMessage `json:"settings,omitempty"`
}

type KeyDownPayload struct {
	Settings         json.RawMessage `json:"settings,omitempty"`
	Coordinates      Coordinates     `json:"coordinates,omitempty"`
	State            int             `json:"state,omitempty"`
	UserDesiredState int             `json:"userDesiredState,omitempty"`
	IsInMultiAction  bool            `json:"isInMultiAction,omitempty"`
}

type KeyUpPayload struct {
	Settings         json.RawMessage `json:"settings,omitempty"`
	Coordinates      Coordinates     `json:"coordinates,omitempty"`
	State            int             `json:"state,omitempty"`
	UserDesiredState int             `json:"userDesiredState,omitempty"`
	IsInMultiAction  bool            `json:"isInMultiAction,omitempty"`
}

type WillAppearPayload struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	State           int             `json:"state,omitempty"`
	IsInMultiAction bool            `json:"isInMultiAction,omitempty"`
}

type WillDisappearPayload struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	State           int             `json:"state,omitempty"`
	IsInMultiAction bool            `json:"isInMultiAction,omitempty"`
}

type TitleParametersDidChangePayload struct {
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

type ApplicationDidLaunchPayload struct {
	Application string `json:"application,omitempty"`
}

type ApplicationDidTerminatePayload struct {
	Application string `json:"application,omitempty"`
}
