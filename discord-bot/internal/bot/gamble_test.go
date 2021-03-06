package bot

import (
	"context"
	"os"
	"testing"

	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

var (
	mes = &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "messageid",
			Content:   "!ping",
			GuildID:   "1",
			ChannelID: "channelid",
			Author: &discordgo.User{
				Username: "userid",
				ID:       "5",
			},
		},
	}
	ctx = context.Background()
)

func TestMain(m *testing.M) {
	DBInt = &db.MockDiscordDBInterface{}
	code := m.Run()
	os.Exit(code)
}

func TestSaveGamble(t *testing.T) {
	type args struct {
		user   model.User
		amount int
		winner bool
	}

	tests := []struct {
		name        string
		args        args
		exceptError bool
	}{
		{
			name: "No db error goes straight through",
			args: args{
				user:   model.User{},
				amount: 5,
				winner: true,
			},
			exceptError: false,
		},
		{
			name: "db error gets caught though",
			args: args{
				user:   model.User{},
				amount: 1,
				winner: true,
			},
			exceptError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := saveGamble(ctx, tt.args.user, tt.args.amount, tt.args.winner)
			if tt.args.amount == 1 {
				assert.NotNil(t, result)
			} else {
				assert.Nil(t, result)
			}
		})
	}
}

func TestCalulatePointsLessThanAll(t *testing.T) { //nolint
	type args struct {
		user        model.User
		amountParam string
		winner      bool
	}

	tests := []struct {
		name        string
		args        args
		want        int
		ExpectError bool
	}{
		{
			name: "Standard winner ok",
			args: args{
				user: model.User{
					UserID:  "5",
					GuildID: "5",
					Points:  20,
				},
				amountParam: "10",
				winner:      true,
			},
			want:        30,
			ExpectError: false,
		},
		{
			name: "Standard loser ok",
			args: args{
				user: model.User{
					UserID:  "5",
					GuildID: "5",
					Points:  20,
				},
				amountParam: "10",
				winner:      false,
			},
			want:        10,
			ExpectError: false,
		},
		{
			name: "Invalid gamble amount",
			args: args{
				user: model.User{
					UserID:  "5",
					GuildID: "5",
					Points:  20,
				},
				amountParam: "-1",
				winner:      false,
			},
			want:        0,
			ExpectError: true,
		},
		{
			name: "Invalid string amount",
			args: args{
				user: model.User{
					UserID:  "5",
					GuildID: "5",
					Points:  20,
				},
				amountParam: "sad",
				winner:      false,
			},
			want:        0,
			ExpectError: true,
		},
		{
			name: "Gamble more points than own",
			args: args{
				user: model.User{
					UserID:  "5",
					GuildID: "5",
					Points:  20,
				},
				amountParam: "30",
				winner:      false,
			},
			want:        0,
			ExpectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calulatePointsLessThanAll(ctx, tt.args.user, tt.args.amountParam, tt.args.winner)
			if tt.ExpectError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalulatePointsAll(t *testing.T) {
	type args struct {
		user   model.User
		winner bool
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Standard winner",
			args: args{
				user: model.User{
					UserID:  "5",
					GuildID: "5",
					Points:  20,
				},
				winner: true,
			},
			want: 40,
		},
		{
			name: "Standard loser ok",
			args: args{
				user: model.User{
					UserID:  "5",
					GuildID: "5",
					Points:  20,
				},
				winner: false,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calulatePointsAll(ctx, tt.args.user, tt.args.winner)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCheckGambleIsSane(t *testing.T) {
	type args struct {
		m *discordgo.MessageCreate
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Standard ok",
			args: args{
				m: mes,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := checkGambleIsSane(ctx, tt.args.m)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
