package crawler

import (
	"github.com/nnhuyhoang/discord_bot/pkg/config"
	"github.com/nnhuyhoang/discord_bot/pkg/consts"
	"github.com/nnhuyhoang/discord_bot/pkg/logger"
	"github.com/nnhuyhoang/discord_bot/pkg/service/discord"

	"github.com/nnhuyhoang/discord_bot/pkg/util"
)

type CrawlerWorker struct {
	Schedule string
	CronWork func()
}

func NewCrawlerClient(cfg config.Config, l logger.Log, discordAction discord.Discord) *CrawlerWorker {
	return &CrawlerWorker{
		Schedule: "0 30 8 * * *",
		CronWork: func() {
			// twitterClient := twitter.NewTwitterClient(cfg, l)
			l.Info("Start crawling from %s", consts.CoronaIndexPage)
			nations := util.GetCoronaIndex()
			message := util.MakeEmbedCoronaData(nations)
			discordAction.SendMessage(cfg.DiscordChannelId, message)
		},
	}
}
