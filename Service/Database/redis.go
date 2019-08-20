package Database

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
)

var Redis *Conn

type Conn struct {
	Client *redis.Client
}

func InitRedis() {
	r := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("Component.Redis.Addr"),
		Password: viper.GetString("Component.Redis.Password"),
		DB:       viper.GetInt("Component.Redis.Database"),
	})
	_, err := r.Ping().Result()
	if err != nil {
		log.Panic("redis init error")
	}
	Redis = &Conn{Client:r,}
}

func (c *Conn) Get()  {

}

func (c *Conn) Set()  {

}

func (c *Conn) SetNx()  {

}
func (c *Conn) Delete()  {

}
