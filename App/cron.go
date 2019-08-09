package App

import (
	"gin/Console"
	"github.com/robfig/cron"
)

func initCron()  {
	go func() {
		c := cron.New()


		//新增一个脚本运行
		_ = c.AddFunc("* * * * *", Console.Test1)


		c.Start()
	}()
}
