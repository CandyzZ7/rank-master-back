package crontask

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func TestCron(t *testing.T) {
	c := cron.New()

	entryID, err := c.AddFunc("@every 5s", func() {
		fmt.Println(time.Now(), "run every 5 second")
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(entryID)
	id, err := c.AddFunc("4 * * * ?", func() {
		fmt.Println(time.Now(), "run minute 4 every hour")
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(id)
	c.Start()
	select {}
}
