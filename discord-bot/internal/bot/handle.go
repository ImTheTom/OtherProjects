package bot

import (
	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var (
	Sess  *discordgo.Session
	DBInt db.DiscordDBInterface
)

func Init() {
	logrus.Info("Initialising bot and commands")

	DBInt = db.GetDatabaseInterface()

	botToken := config.GetConfig().BotToken

	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		logrus.Fatalf("Discord session start failed, restarting... %v", err)
	}

	session.AddHandler(standardChatMessages)

	session.AddHandler(imageUploadMessage)

	session.AddHandler(gambleInteractions)

	// Maybe revisit if needbe
	session.Identify.Intents = discordgo.IntentsAll

	err = session.Open()
	if err != nil {
		logrus.Fatalf("Discord failed to open the connection, restarting... %v", err)
	}

	logrus.Info("Bot is now running...")

	Sess = session
}

func GetSession() *discordgo.Session {
	if Sess != nil {
		return Sess
	}

	Init()

	return Sess
}

func CloseBot() {
	if Sess != nil {
		Sess.Close()
	}
}

func communicateStandardMessage(s *discordgo.Session, m *discordgo.MessageCreate, message string) {
	if _, err := s.ChannelMessageSend(m.ChannelID, message); err != nil {
		logrus.Errorf("Failed to send message %v", err)

		return
	}

	logrus.WithFields(logrus.Fields{
		"id":      m.ID,
		"content": message,
	}).Info("Command was handled")
}

func logMessage(m *discordgo.MessageCreate) {
	logrus.WithFields(logrus.Fields{
		"id":            m.ID,
		"channel_id":    m.ChannelID,
		"guild_id":      m.GuildID,
		"content":       m.Content,
		"user_id":       m.Author.ID,
		"user_username": m.Author.Username,
	}).Info("Command was called")
}
