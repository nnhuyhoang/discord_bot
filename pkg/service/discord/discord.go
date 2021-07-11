package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/nnhuyhoang/discord_bot/pkg/config"
	"github.com/nnhuyhoang/discord_bot/pkg/logger"
	"github.com/nnhuyhoang/discord_bot/pkg/model"
)

type BotSession struct {
	Session *discordgo.Session
	Channel string
	Log     logger.Log
}

type Discord interface {
	AddTask(callback interface{}) error
	SendMessage(channelId string, message interface{}) error
}

func NewBotDiscord(cfg config.Config, l logger.Log) *BotSession {
	dg, err := discordgo.New("Bot " + cfg.DiscordBotToken)
	if err != nil {
		log.Fatal(err)
	}
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages
	return &BotSession{
		Session: dg,
		Channel: cfg.DiscordChannelId,
		Log:     l,
	}
}

func (b *BotSession) SendMessage(channelId string, message interface{}) error {

	switch message := message.(type) {
	case string:
		_, err := b.Session.ChannelMessageSend(channelId, message)
		if err != nil {
			b.Log.Error("[BotSession.SendMessage] b.Session.ChannelMessageSend(channelId, message): ", err)
			return err
		}
	case discordgo.MessageEmbed:
		_, err := b.Session.ChannelMessageSendEmbed(channelId, &message)
		if err != nil {
			b.Log.Error("[BotSession.SendMessage] b.Session.ChannelMessageSendEmbed(channelId, &message): ", err)
			return err
		}
	default:
		b.Log.Error("[BotSession.SendMessage] Unknown type message")
		return &model.InternalError{}
	}
	return nil
}

func (b *BotSession) AddTask(callback interface{}) error {
	switch callback := callback.(type) {
	case func(*discordgo.Session, *discordgo.MessageCreate):
		b.Session.AddHandler(callback)
	default:
		b.Log.Error("[BotSession.AddTask] invliad type callback")
		return &model.InternalError{}
	}
	return nil
}
