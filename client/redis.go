package client

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/configor"
	"cradle/walle/common"
)
var Config common.Config

func LoadConfig() {
	configor.Load(&Config, "conf/walle.yaml")
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

/* only once
func SaveKeyINRedis()  {
	GetRedisClient().Set("corpid","ww63adb0909f572435", 0)
	GetRedisClient().Set("corpsecret","_k1tl1qmsntCBHVITGrKqbSopi7E18Q8joYjeD62_K0",0)
	GetRedisClient().Set("private_token","1u-v5yZEuZsj18yhg7ai",0)
}
func main() {
	SaveKeyINRedis()
}
*/