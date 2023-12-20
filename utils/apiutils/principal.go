package apiutils

import "context"

type contextKey int

const (
	contextKeyPrincipalAddressValue contextKey = 1
	contextKeyBatchValue            contextKey = 2
	contextKeyBotID                 contextKey = 3
)

func GetAddress(ctx context.Context) string {
	return ctx.Value(contextKeyPrincipalAddressValue).(string)
}

func SetAddress(ctx context.Context, addr string) context.Context {
	return context.WithValue(ctx, contextKeyPrincipalAddressValue, addr)
}

func GetBatch(ctx context.Context) string {
	return ctx.Value(contextKeyBatchValue).(string)
}

func SetBatch(ctx context.Context, batch string) context.Context {
	return context.WithValue(ctx, contextKeyBatchValue, batch)
}

func GetBotID(ctx context.Context) string {
	return ctx.Value(contextKeyBotID).(string)
}

func SetBotID(ctx context.Context, addr string) context.Context {
	return context.WithValue(ctx, contextKeyBotID, addr)
}
