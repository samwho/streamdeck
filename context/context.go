package context

import (
	"context"
)

type keyType int

const (
	contextKey keyType = iota
	deviceKey
	actionKey
)

func Context(ctx context.Context) string {
	return get(ctx, contextKey)
}

func WithContext(ctx context.Context, streamdeckContext string) context.Context {
	return context.WithValue(ctx, contextKey, streamdeckContext)
}

func Device(ctx context.Context) string {
	return get(ctx, deviceKey)
}

func WithDevice(ctx context.Context, streamdeckDevice string) context.Context {
	return context.WithValue(ctx, deviceKey, streamdeckDevice)
}

func Action(ctx context.Context) string {
	return get(ctx, actionKey)
}

func WithAction(ctx context.Context, streamdeckAction string) context.Context {
	return context.WithValue(ctx, actionKey, streamdeckAction)
}

func get(ctx context.Context, key keyType) string {
	if ctx == nil {
		return ""
	}

	val := ctx.Value(key)
	if val == nil {
		return ""
	}

	valStr, ok := val.(string)
	if !ok {
		panic("found non-string in context")
	}

	return valStr
}
