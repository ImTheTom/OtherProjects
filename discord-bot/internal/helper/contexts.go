package helper

import (
	"context"
	"time"
)

const numberOfMilliseconds = 500

func CreateContextWithTimeout() (context.Context, func()) {
	ctxBg := context.Background()

	ctx, cancel := context.WithTimeout(ctxBg, time.Millisecond*numberOfMilliseconds)

	return ctx, cancel
}
