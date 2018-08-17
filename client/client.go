package client

import (
	"github.com/parnurzeal/gorequest"
	"autoDelUser/common"
	"encoding/json"
	"time"
	"strconv"
)

const(
 corpid string ="ww63adb0909f572435"
 corpsecret string = "_k1tl1qmsntCBHVITGrKqbSopi7E18Q8joYjeD62_K0"
 getQyTokenUrl string = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid="
 getQyUsersUrl string = "https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token="
 getGitlabUsersUrl string = "https://gitlab.wallstcn.com/api/v4/users"
 private_token string = "1u-v5yZEuZsj18yhg7ai"
)

var qyToken common.QyToken
var qyUser common.QyUser
var qyUserItem common.QyUserItem
//var gitlabUser common.GitlabUser

var GitlabEmailMap = make(map[string]common.GitlabUser)

var QyEmailMap = make(map[string]common.QyUserItem)
//var Config common.Config

//func LoadConfig() {
//	configor.Load(&Config, "conf/autoDelUser.yaml")
//}

/* get token and Cache to redis*/
func GetQyToken() string {
	//LoadConfig()
	/* read redis cache */
	//client := redis.NewClient(&redis.Options{
	//	//Addr:     "localhost:6379",
	//	//Password: "", // no password set
	//	//DB:       0,  // use default DB
	//	Addr:     Config.Redis.Addr,
	//	Password: Config.Redis.Password,
	//	DB:       Config.Redis.Db,
	//})
	val, err := GetRedisClient().Get("qytoken").Result()
	/* save data into redis */
	if val == "" && err != nil {
		request := gorequest.New()
		resp, body, errs := request.Get(getQyTokenUrl + corpid + "&corpsecret=" + corpsecret).End()
		if resp.StatusCode != 200 && errs != nil {
			panic(err)
		}
		json.Unmarshal([]byte(body), &qyToken)
		err := GetRedisClient().Set("qytoken", qyToken.AccessToken, 7100*time.Second).Err()
		val = qyToken.AccessToken
		if err != nil {
			panic(err)
		}
	}
	return val

}

/* get qiye weixin Users*/
func GetQyUsers() {
	request := gorequest.New()
	resp, body, errs := request.Get(getQyUsersUrl + GetQyToken() + "&department_id=1&fetch_child=1").End()
	if resp.StatusCode != 200 && errs != nil {
		panic(errs)
	}
	qyUser.UserList = make([]common.QyUserItem, 10)
	json.Unmarshal([]byte(body), &qyUser)

	for _, val := range qyUser.UserList {
		QyEmailMap[val.Email] = val
	}

}

/* get gitlab all users */

func GetGitlabUsers() {
	request := gorequest.New()
	gitlabUser := make([]common.GitlabUser, 10)
	for page := 1; page < 7; page ++ {
		resp, body, errs := request.Get(getGitlabUsersUrl + "?" + "private_token=" + private_token + "&" + "page=" + strconv.Itoa(page)).End()
		if resp.StatusCode != 200 && errs != nil {
			panic(errs)
		}
		json.Unmarshal([]byte(body), &gitlabUser)
		for _, val := range gitlabUser {
			GitlabEmailMap[val.Email] = val
		}

	}

}

/* block gitlab users */
func BlockGitlabUsers(userId int64) string {
	request := gorequest.New()
	resp, body, errs := request.Post(getGitlabUsersUrl + "/" + strconv.FormatInt(userId, 10) + "/block" + "?private_token=" + private_token).End()
	if resp.StatusCode != 200 && errs != nil {
		panic(errs)
	}
	return body
}
