package client

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/configor"
	"autoDelUser/common"
)
var Config common.Config

func LoadConfig() {
	configor.Load(&Config, "conf/autoDelUser.yaml")
}

func GetRedisClient() *redis.Client{
	LoadConfig()
	/* read redis cache */
	client := redis.NewClient(&redis.Options{
		//Addr:     "localhost:6379",
		//Password: "", // no password set
		//DB:       0,  // use default DB
		Addr:     Config.Redis.Addr,
		Password: Config.Redis.Password,
		DB:       Config.Redis.Db,
	})
	return client
}
