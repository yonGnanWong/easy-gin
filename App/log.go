package App

import (
	"log"
	"os"
	"time"
)

//log日志文件的 io writer
var LogWriter *os.File

//初始化log配置
func InitLog()  {
	file := "./logs/" + time.Now().Format("2019-7-02") + ".log"
	logfile, _ := os.OpenFile(file,os.O_RDWR| os.O_CREATE| os.O_APPEND, 0755)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logfile)
	LogWriter = logfile
}
