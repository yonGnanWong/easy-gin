package Database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

//申明全局变量Db
var Db *xorm.Engine

//初始化方法
func InitDb() {
	//初始化数据库配置参数
	DataType := viper.GetString("Component.Database.DriverName")
	Addr := viper.GetString("Component.Database.Addr")
	MaxOpenConns := viper.GetInt("Component.Database.MaxOpenConns")
	MaxIdleConns := viper.GetInt("Component.Database.MaxIdleConns")

	db, err := xorm.NewEngine(DataType, Addr)
	if err != nil {
		log.Panic("database error ", err)
	}
	if db == nil {
		log.Panic("database error ")
	}

	if viper.GetBool("Debug") {
		//则会在控制台打印出生成的SQL语句
		db.ShowSQL(true)
		//会在控制台打印调试及以上的信息
		db.Logger().SetLevel(core.LOG_DEBUG)
	}

	//将sql写入log
	//db.SetLogger(xorm.NewSimpleLogger(LogWriter))
	//设置db连接池最大链接数和最大打开链接数
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)

	Db = db
}
