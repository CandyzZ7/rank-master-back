package crontask

import (
	"github.com/robfig/cron/v3"
)

type Cron struct {
	cron *cron.Cron
}

func NewCron() *Cron {
	return &Cron{cron.New(cron.WithChain(cron.SkipIfStillRunning(&CLog{})))}
}

func (c *Cron) StartTask() {
	c.cron.Start()
}

func (c *Cron) StopTask() {
	c.cron.Stop()
}

func (c *Cron) AddTask(spec string, task cron.Job) (cron.EntryID, error) {
	entryID, err := c.cron.AddJob(spec, task)
	if err != nil {
		return 0, err
	}

	return entryID, nil
}

func (c *Cron) WaitTask() {
	select {}
}
