package cronjob

import (
	"time"

	"github.com/nnhuyhoang/discord_bot/pkg/consts"
	"github.com/robfig/cron"
)

type cronManager struct {
	manager *cron.Cron
}

type CronManager interface {
	AddTask(scheduleStr string, fn func()) error
	Start() func()
}

func NewCronManager() CronManager {
	loc, err := time.LoadLocation(consts.DefaultTimeZone)
	if err != nil {
		loc = time.Local
	}
	c := cron.NewWithLocation(loc)
	return &cronManager{
		manager: c,
	}

}

func (cm *cronManager) AddTask(scheduleStr string, fn func()) error {
	return cm.manager.AddFunc(scheduleStr, fn)
}
func (cm *cronManager) Start() func() {
	cm.manager.Start()
	return cm.manager.Stop
}
