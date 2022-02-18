// GENERATED CODE DO NOT EDIT
// DO NOT EDIT
// RE RUN ./scripts/regenerate.sh

package bot

import (
	"fmt"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

const helpMessage = `The bot has many commands to interact with here is the list:

flip - Flips a coin.
hello - Prints a welcome message.
ping - General ping command.
rap - The famous rap command.
ladder - Prints out a leaderboard of the top 10 users in the channel.
birth - Displays a fun message about when the bot was first created.
louis - Prints out a message releating to Louis.
points - Will print out how many points the user has.
gamble - Can gamble your points. Used by providing an additional paramater.
image - Sends a fun image. Send Tom suggestions.
joined - Displays a fun message about when you joined the guild.`

var _fullHelpMessage string

func helpInteraction(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix := config.GetConfig().Prefix

	if _fullHelpMessage == "" {
		_fullHelpMessage = fmt.Sprintf("%s\n\nRemember to use your prefix - %s", helpMessage, prefix)
	}

	switch m.Content {
	case fmt.Sprintf("%s%s", prefix, "help"):
		communicateStandardMessage(s, m, _fullHelpMessage)
	}
}
