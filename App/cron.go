package App

import (
	"fmt"
	"github.com/robfig/cron"
)

func initCron()  {
	go func() {
		c := cron.New()
		_ = c.AddFunc("*/1 * * * *", func() {
			i := 0
			i++
			fmt.Printf("cron running %d",i)
		})
		c.Start()
	}()
}
