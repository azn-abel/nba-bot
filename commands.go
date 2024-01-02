package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandError struct {
	Message string
}

func (cmdErr CommandError) Error() string {
	return cmdErr.Message
}

var APIBaseURL string = "https://site.api.espn.com/apis/site/v2/sports/basketball/nba/"

func Scoreboard(words []string) (*discordgo.MessageEmbed, error) {
	endpointURL := APIBaseURL + "scoreboard"
	response, err := http.Get(endpointURL)
	if err != nil {
		fmt.Println("endpoint err")
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var scoreboardData ScoreboardStruct

	err = json.Unmarshal(body, &scoreboardData)
	if err != nil {
		return nil, err
	}

	fmt.Println(scoreboardData)

	embed := &discordgo.MessageEmbed{
		Title:       "Woohoo!",
		Description: "Successfully unmarshalled json",
	}

	return embed, nil
}

func Team(words []string) (*discordgo.MessageEmbed, error) {

	if len(words) < 2 {
		return nil, CommandError{
			Message: "Not enough args",
		}
	}

	teamName := strings.ToLower(words[1])

	value, exists := TeamIDByName[teamName]
	if !exists {
		return nil, CommandError{
			Message: "Invalid team name",
		}
	}

	endpointURL := APIBaseURL + "teams/" + fmt.Sprintf("%d", value)
	response, err := http.Get(endpointURL)
	if err != nil {
		fmt.Println("endpoint err")
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var teamData TeamResponse

	err = json.Unmarshal(body, &teamData)
	if err != nil {
		return nil, err
	}

	team := teamData.Team
	color, err := strconv.ParseUint(team.Color, 16, 32)
	if err != nil {
		return nil, err
	}

	thumbnail := &discordgo.MessageEmbedThumbnail{
		URL: team.Logos[0].URL,
	}
	overallRecordField := &discordgo.MessageEmbedField{
		Name:   "Overall",
		Value:  team.Record.Items[0].Summary,
		Inline: true,
	}
	homeRecordField := &discordgo.MessageEmbedField{
		Name:   "Home",
		Value:  team.Record.Items[1].Summary,
		Inline: true,
	}
	awayRecordField := &discordgo.MessageEmbedField{
		Name:   "Away",
		Value:  team.Record.Items[2].Summary,
		Inline: true,
	}
	overallPpgField := &discordgo.MessageEmbedField{
		Name:   "PPG",
		Value:  floatToString(team.Record.Items[0].Stats[3].Value),
		Inline: true,
	}
	overallPpgAgainstField := &discordgo.MessageEmbedField{
		Name:   "PPG Against",
		Value:  floatToString(team.Record.Items[0].Stats[2].Value),
		Inline: true,
	}
	ppgDifferentialField := &discordgo.MessageEmbedField{
		Name:   "Differential",
		Value:  floatToString(team.Record.Items[0].Stats[4].Value),
		Inline: true,
	}

	fields := []*discordgo.MessageEmbedField{
		overallRecordField,
		homeRecordField,
		awayRecordField,
		overallPpgField,
		overallPpgAgainstField,
		ppgDifferentialField,
	}

	streak := team.Record.Items[0].Stats[15].Value
	var streakString string
	if streak < 0 {
		streakString = "Lost " + fmt.Sprintf("%d", int(math.Trunc(math.Abs(streak)))) + " in a row"
	} else {
		streakString = "Won " + fmt.Sprintf("%d", int(math.Trunc(math.Abs(streak)))) + " in a row"
	}

	description := fmt.Sprintf("**Next Game (%s):**\n", team.NextEvent[0].Date[:10]) + team.NextEvent[0].Name + "\n\n"
	description += "**Streak:**\n" + streakString + "\n\n** **"

	embed := &discordgo.MessageEmbed{
		Title:       team.Name,
		Description: description,
		URL:         team.Links[0].URL,
		Thumbnail:   thumbnail,
		Fields:      fields,
		Color:       int(color),
	}

	return embed, nil
}

func floatToString(num float64) string {
	return fmt.Sprintf("%.2f", num)
}
