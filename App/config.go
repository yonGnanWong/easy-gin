package App

import (
	"github.com/spf13/viper"
	"log"
)

var Config ConfigParam

//配置参数
type ConfigParam struct{
	DEBUG bool
	ENV string
	Port string
	RunMode string
}

//初始化config数据
func initConfig() {
	//viper 读取配置文件参数
	viper.AddConfigPath("Config")
	//读取dev配置文件.如果设计配置文件的切换请在提交前修改
	//TODO:这里是手动切换配置文件,没有想到更好的方法
	viper.SetConfigName("dev")
	err := viper.ReadInConfig()
	if err != nil{
		log.Panic(err)
	}

	//初始化全局config结构体
	Config.DEBUG = viper.GetBool("Debug")
	Config.ENV = viper.GetString("Env")
	Config.Port = viper.GetString("Port")
	Config.RunMode = viper.GetString("RunMode")
}
