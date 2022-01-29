package helper

import (
	"context"
	"time"
)

const timeoutMs = 150

func CreateContextWithTimeout() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeoutMs)*time.Millisecond)

	return ctx
}
