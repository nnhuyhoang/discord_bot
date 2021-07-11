package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nnhuyhoang/discord_bot/pkg/consts"
	"github.com/nnhuyhoang/discord_bot/pkg/model"
)

func MakeEmbedCoronaData(countries []model.Country) discordgo.MessageEmbed {
	var messageMarkdowns []string
	currentTime := time.Now()
	//Format MM-DD-YYYY
	currentDate := currentTime.Format("Jan-02")
	currentDateStr := fmt.Sprintf("[%s]", currentDate)
	messageMarkdowns = append(messageMarkdowns, currentDateStr)
	for _, country := range countries {
		messageMarkdowns = append(messageMarkdowns, MakeMarkDownData(country))
	}
	finalMessage := strings.Join(messageMarkdowns, "\n")
	return discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Description: finalMessage,
	}
}

func MakeEmbedCoronaSingleData(country model.Country) discordgo.MessageEmbed {
	messageMarkdown := MakeMarkDownData(country)
	return discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Description: messageMarkdown,
	}
}

func MakeMarkDownData(country model.Country) string {
	flag := consts.FlagEmojiMapper[country.Name]
	return fmt.Sprintf(`%s %s **Total Case**: %s (%s), **Total Death**: %s (%s)`, flag, country.Name, country.TotalCase, country.NewCase, country.TotalDeath, country.NewDeath)
}
