package streamdeck

type Action struct {
	uuid     string
	handlers map[string][]EventHandler
}

func (action *Action) RegisterHandler(eventName string, handler EventHandler) {
	action.handlers[eventName] = append(action.handlers[eventName], handler)
}
