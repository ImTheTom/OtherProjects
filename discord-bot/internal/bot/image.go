package bot

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

const (
	imagePath    = "./assets/images/"
	imageCommand = "image"
)

func ImageUploadMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	files, err := ioutil.ReadDir(imagePath)
	if err != nil {
		fmt.Printf("Failed %v\n", err)
		return
	}

	index := rand.Intn(len(files))

	reader, err := os.Open(fmt.Sprintf("%s/%s", imagePath, files[index].Name()))
	if err != nil {
		fmt.Printf("Failed %v\n", err)
		return
	}

	if _, err := s.ChannelFileSend(m.ChannelID, files[index].Name(), reader); err != nil {
		fmt.Printf("failed to send message %v\n", err)
	}
}
