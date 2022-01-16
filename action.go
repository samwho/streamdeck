package streamdeck

import (
	"context"

	sdcontext "github.com/samwho/streamdeck/context"
)

type Action struct {
	uuid     string
	handlers map[string][]EventHandler
	contexts map[string]context.Context
}

func newAction(uuid string) *Action {
	action := &Action{
		uuid:     uuid,
		handlers: make(map[string][]EventHandler),
		contexts: make(map[string]context.Context),
	}

	action.RegisterHandler(WillAppear, func(ctx context.Context, client *Client, event Event) error {
		action.addContext(ctx)
		return nil
	})

	action.RegisterHandler(WillDisappear, func(ctx context.Context, client *Client, event Event) error {
		action.removeContext(ctx)
		return nil
	})

	return action
}

func (action *Action) RegisterHandler(eventName string, handler EventHandler) {
	action.handlers[eventName] = append(action.handlers[eventName], handler)
}

func (action *Action) Contexts() []context.Context {
	cs := make([]context.Context, len(action.contexts))
	for _, c := range action.contexts {
		cs = append(cs, c)
	}
	return cs
}

func (action *Action) addContext(ctx context.Context) {
	if sdcontext.Context(ctx) == "" {
		panic("passed non-streamdeck context to addContext")
	}

	action.contexts[sdcontext.Context(ctx)] = ctx
}

func (action *Action) removeContext(ctx context.Context) {
	if sdcontext.Context(ctx) == "" {
		panic("passed non-streamdeck context to addContext")
	}

	delete(action.contexts, sdcontext.Context(ctx))
}
