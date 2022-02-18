package bot

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

const imagePath = "./assets/images/"

// Sends a fun image. Send Tom suggestions.
const imageCommand = "image"

func imageUploadMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix := config.GetConfig().Prefix

	if m.Content == fmt.Sprintf("%s%s", prefix, imageCommand) {
		imageMessage(s, m)
	}
}

func imageMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	logMessage(m)

	files, err := ioutil.ReadDir(imagePath)
	if err != nil {
		logrus.Errorf("Read Dir Failed %v", err)

		return
	}

	index := rand.Intn(len(files))

	reader, err := os.Open(fmt.Sprintf("%s/%s", imagePath, files[index].Name()))
	if err != nil {
		logrus.Errorf("OS Open Failed %v", err)

		return
	}

	if _, err := s.ChannelFileSend(m.ChannelID, files[index].Name(), reader); err != nil {
		logrus.Errorf("failed to send file %v", err)

		return
	}

	logrus.WithFields(logrus.Fields{
		"id":      m.ID,
		"content": files[index].Name(),
	}).Info("Command was handled")
}
