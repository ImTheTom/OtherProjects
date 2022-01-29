package helper

import (
	"context"
)

const timeoutMs = 150

func CreateContextWithTimeout() context.Context {
	return context.TODO()
}
