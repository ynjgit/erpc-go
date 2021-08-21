package protocol

import (
	"context"
)

type contextKey string

const (
	ctxKey contextKey = "erpc-ctx"
)

// PbasCtx ...
type PbasCtx struct {
	ServerRPCName string

	ClientRPCName string

	Logger interface{}
}

// GetCtx ...
func GetCtx(ctx context.Context) *PbasCtx {
	v := ctx.Value(ctxKey)
	if c, ok := v.(*PbasCtx); ok {
		return c
	}

	return &PbasCtx{}
}

// SetCtx ...
func SetCtx(ctx context.Context, erpcCtx *PbasCtx) context.Context {
	return context.WithValue(ctx, ctxKey, erpcCtx)
}
