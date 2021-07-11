package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nnhuyhoang/discord_bot/pkg/util"
)

func (b *BotSession) EventCreateMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ChannelID != b.Channel {
		return
	}

	if strings.HasPrefix(m.Content, ".") {
		if strings.HasPrefix(m.Content, ".s") {
			nationName := strings.Split(m.Content, ".s")
			if strings.TrimSpace(nationName[1]) != "" {
				nation := util.GetCoronaIndexByCountryName(strings.TrimSpace(nationName[1]))
				message := util.MakeEmbedCoronaSingleData(nation)
				b.SendMessage(m.ChannelID, message)
			} else {
				b.SendMessage(m.ChannelID, "invalid parameter")
			}
		} else {
			b.SendMessage(m.ChannelID, "invalid command")
		}
	}
}
