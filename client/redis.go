package client

import (
	"cradle/walle/common"
	"github.com/go-redis/redis"
	"github.com/jinzhu/configor"
	"os"
)

var Config common.Config

func LoadConfig() {
	env := os.Getenv("CONFIG_ENV")
	if env == "prod"{
		err := configor.Load(&Config, "/conf/walle.yaml")
		if err != nil {
			panic(err)
		}
	}else  {
		err := configor.Load(&Config, "/conf/walle_test.yaml")
		if err != nil {
			panic(err)
		}
	}
}

func GetRedisClient() *redis.Client {
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
	GetRedisClient().Set("corpid","*", 0)
	GetRedisClient().Set("corpsecret","*",0)
	GetRedisClient().Set("private_token","*",0)
}
func main() {
	SaveKeyINRedis()
}
*/
