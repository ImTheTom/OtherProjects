package bot_test

import (
	"testing"

	"github.com/ImTheTom/OtherProjects/discord-bot/internal/bot"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/stretchr/testify/assert"
)

func TestSaveGamble(t *testing.T) {
	type args struct {
		user   model.User
		amount int
		winner bool
	}

	tests := []struct {
		name                 string
		args                 args
		dbInteractionReturns error
		exceptError          bool
	}{
		{
			name: "No db error goes straight through",
			args: args{
				user:   model.User{},
				amount: 5,
				winner: true,
			},
			dbInteractionReturns: nil,
			exceptError:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot.DBInt = db.MockDiscordDBInterface{}
			result := bot.SaveGamble(tt.args.user, tt.args.amount, tt.args.winner)
			assert.Nil(t, result)
		})
	}
}
