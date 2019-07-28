package App

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"log"
)

//申明全局变量Db
var Db DbParam

//数据库参数
type DbParam struct{
	DataType string
	Addr string
	MaxOpenConns int
	MaxIdleConns int
	Engine *xorm.Engine
}

//初始化方法
func InitDb() {
	//初始化数据库配置参数
	Db.DataType = viper.GetString("Component.Database.DriverName")
	Db.Addr = viper.GetString("Component.Database.Addr")
	Db.MaxOpenConns = viper.GetInt("Component.Database.MaxOpenConns")
	Db.MaxIdleConns = viper.GetInt("Component.Database.MaxIdleConns")

	db, err := xorm.NewEngine(Db.DataType,Db.Addr)
	if err != nil {
		log.Panic("database error ", err)
	}
	if db == nil {
		log.Panic("database error ")
	}

	if Config.DEBUG {
		//则会在控制台打印出生成的SQL语句
		db.ShowSQL(true)
		//会在控制台打印调试及以上的信息
		db.Logger().SetLevel(core.LOG_DEBUG)
	}

	//将sql写入log
	db.SetLogger(xorm.NewSimpleLogger(LogWriter))

	//设置db连接池最大链接数和最大打开链接数
	Db.Engine = db
	Db.Engine.SetMaxIdleConns(Db.MaxIdleConns)
	Db.Engine.SetMaxOpenConns(Db.MaxOpenConns)
}
