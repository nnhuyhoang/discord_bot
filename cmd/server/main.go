package main

import (
	"log"

	"github.com/nnhuyhoang/discord_bot/pkg/config"
	cronjob "github.com/nnhuyhoang/discord_bot/pkg/cron"
	"github.com/nnhuyhoang/discord_bot/pkg/logger"
	"github.com/nnhuyhoang/discord_bot/pkg/service/crawler"
	"github.com/nnhuyhoang/discord_bot/pkg/service/discord"
)

func main() {
	cls := config.DefaultConfigLoaders()
	cfg := config.LoadConfig(cls)
	l := initLog(cfg)

	bot := initDiscordBot(cfg, l)
	defer bot.Session.Close()

	cron := initCronjJob(cfg, l, bot)
	cStop := cron.Start()
	defer cStop()

	select {}

}

func initLog(cfg config.Config) logger.Log {
	return logger.NewJSONLogger(
		logger.WithServiceName(cfg.ServiceName),
		logger.WithHostName(cfg.BaseURL),
	)
}

func initDiscordBot(cfg config.Config, l logger.Log) *discord.BotSession {
	bot := discord.NewBotDiscord(cfg, l)
	bot.AddTask(bot.EventCreateMessage)
	err := bot.Session.Open()
	if err != nil {
		log.Fatal("error open")
	}
	return bot
}

func initCronjJob(cfg config.Config, l logger.Log, bot *discord.BotSession) cronjob.CronManager {
	crawler := crawler.NewCrawlerClient(cfg, l, bot)
	cron := cronjob.NewCronManager()

	cron.AddTask(crawler.Schedule, crawler.CronWork)
	return cron
}
