package streamdeck

import "context"

const (
	DidReceiveSettings            = "didReceiveSettings"
	DidReceiveGlobalSettings      = "didReceiveGlobalSettings"
	KeyDown                       = "keyDown"
	KeyUp                         = "keyUp"
	WillAppear                    = "willAppear"
	WillDisappear                 = "willDisappear"
	TitleParametersDidChange      = "titleParametersDidChange"
	DeviceDidConnect              = "deviceDidConnect"
	DeviceDidDisconnect           = "deviceDidDisconnect"
	ApplicationDidLaunch          = "applicationDidLaunch"
	ApplicationDidTerminate       = "applicationDidTerminate"
	SystemDidWakeUp               = "systemDidWakeUp"
	PropertyInspectorDidAppear    = "propertyInspectorDidAppear"
	PropertyInspectorDidDisappear = "propertyInspectorDidDisappear"
	SendToPlugin                  = "sendToPlugin"
	SendToPropertyInspector       = "sendToPropertyInspector"

	SetSettings       = "setSettings"
	GetSettings       = "getSettings"
	SetGlobalSettings = "setGlobalSettings"
	GetGlobalSettings = "getGlobalSettings"
	OpenURL           = "openUrl"
	LogMessage        = "logMessage"
	SetTitle          = "setTitle"
	SetImage          = "setImage"
	ShowAlert         = "showAlert"
	ShowOk            = "showOk"
	SetState          = "setState"
	SwitchToProfile   = "switchToProfile"
)

type contextKeyType int

const (
	contextKey contextKeyType = iota
)

func getContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	return ctx.Value(contextKey).(string)
}

func setContext(ctx context.Context, streamdeckContext string) context.Context {
	return context.WithValue(ctx, contextKey, streamdeckContext)
}
