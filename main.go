package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

var Token = os.Getenv("DISCORD_BOT_TOKEN")
var Prefix string = "nba!"

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsAll

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Gopher is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	fmt.Println("Shutting down...")
	dg.Close()
}

// Runs on every message that the bot can read
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore messages that are too short or don't start with Prefix
	if len(m.Content) < 3 || !strings.HasPrefix(strings.ToLower(m.Content), Prefix) {
		return
	}

	// Break message up into words
	words := strings.Fields(m.Content)

	fmt.Println(words)

	switch strings.ToLower(words[0]) {
	case Command("ping"):
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case Command("pong"):
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	case Command("embed"):
		embed := CreateEmbed()
		// image := CreateImageOnlyEmbed("https://go.dev/doc/gopher/gopher5logo.jpg")
		s.ChannelMessageSendEmbeds(m.ChannelID, []*discordgo.MessageEmbed{embed})
	case Command("docs"):
		s.ChannelMessageSendReply(m.ChannelID, "https://go.dev/doc/", (*m).Reference())
	case Command("nba"):
		embed := CreateNBAEmbed()
		// image1 := CreateImageOnlyEmbed("https://a4.espncdn.com/combiner/i?img=%2Fi%2Fespn%2Fmisc_logos%2F500%2Fnba.png")
		// image2 := CreateImageOnlyEmbed("https://a.espncdn.com/i/teamlogos/nba/500/scoreboard/mil.png")
		image3 := CreateThumbnailOnlyEmbed("https://a.espncdn.com/i/teamlogos/nba/500/scoreboard/ny.png")
		s.ChannelMessageSendEmbeds(m.ChannelID, []*discordgo.MessageEmbed{embed, image3})
	case Command("team"):
		res, err := Team(words)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}

		_, sendErr := s.ChannelMessageSendEmbed(m.ChannelID, res)
		if sendErr != nil {
			s.ChannelMessageSend(m.ChannelID, sendErr.Error())
			return
		}
	case Command("scoreboard"):
		res, err := Scoreboard(words)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		_, sendErr := s.ChannelMessageSendEmbed(m.ChannelID, res)
		if sendErr != nil {
			s.ChannelMessageSend(m.ChannelID, sendErr.Error())
			return
		}
	}

}

func Command(cmd string) string {
	res := Prefix + cmd
	return res
}

func CreateEmbed() *discordgo.MessageEmbed {
	image := &discordgo.MessageEmbedImage{
		URL: "https://pkg.go.dev/static/shared/gopher/package-search-700x300.jpeg",
	}
	footer := &discordgo.MessageEmbedFooter{
		Text: "Me, right now, developing this bot",
	}
	embed := &discordgo.MessageEmbed{
		Title:       "Gopher",
		Description: "This bot is written in... you guessed it, Golang!",
		Color:       0x00add8, // Gopher blue
		Image:       image,
		Footer:      footer,
		URL:         "https://github.com/azn-abel",
	}
	return embed
}

func CreateImageOnlyEmbed(url string) *discordgo.MessageEmbed {
	image := &discordgo.MessageEmbedImage{
		URL: url,
	}
	embed := &discordgo.MessageEmbed{
		Image: image,
		URL:   "https://github.com/azn-abel",
	}
	return embed
}

func CreateThumbnailOnlyEmbed(url string) *discordgo.MessageEmbed {
	thumbnail := &discordgo.MessageEmbedThumbnail{
		URL: url,
	}
	embed := &discordgo.MessageEmbed{
		Thumbnail: thumbnail,
		URL:       "https://github.com/azn-abel",
	}
	return embed
}

func CreateNBAEmbed() *discordgo.MessageEmbed {
	homeScoreField := &discordgo.MessageEmbedField{
		Name:   "Bucks",
		Value:  "108",
		Inline: true,
	}
	awayScoreField := &discordgo.MessageEmbedField{
		Name:   "Knicks",
		Value:  "109",
		Inline: true,
	}
	periodField := &discordgo.MessageEmbedField{
		Name:   "Period",
		Value:  "4",
		Inline: true,
	}
	timeField := &discordgo.MessageEmbedField{
		Name:   "Time left",
		Value:  "2:02",
		Inline: true,
	}
	fields := []*discordgo.MessageEmbedField{
		homeScoreField, awayScoreField, periodField, timeField,
	}
	thumbnail := &discordgo.MessageEmbedThumbnail{
		URL: "https://a.espncdn.com/i/teamlogos/nba/500/scoreboard/mil.png",
	}

	embed := &discordgo.MessageEmbed{
		Title:     "Bucks vs. Knicks",
		Color:     0x00add8,
		URL:       "https://github.com/azn-abel",
		Fields:    fields,
		Thumbnail: thumbnail,
	}
	return embed
}
