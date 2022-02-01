package helper_test

import (
	"context"
	"testing"

	"github.com/ImTheTom/OtherProjects/discord-bot/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestCreateContextWithTimeout(t *testing.T) {
	tests := []struct {
		name string
		want context.Context
	}{
		{
			name: "standard",
			want: context.TODO(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, helper.CreateContextWithTimeout())
		})
	}
}
