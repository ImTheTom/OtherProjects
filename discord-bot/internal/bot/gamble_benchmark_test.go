package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

var mes2 = &discordgo.MessageCreate{
	Message: &discordgo.Message{
		ID:        "messageid",
		Content:   "!ping",
		GuildID:   "1",
		ChannelID: "channelid",
		Author: &discordgo.User{
			Username: "userid",
			ID:       "1",
		},
	},
}

func benchmarkCheckGambleIsSane(m *discordgo.MessageCreate, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = checkGambleIsSane(ctx, m)
	}
}

func BenchmarkCheckGambleIsSaneMes(b *testing.B) {
	benchmarkCheckGambleIsSane(mes, b)
}

func BenchmarkCheckGambleIsSaneMes2(b *testing.B) {
	benchmarkCheckGambleIsSane(mes2, b)
}
