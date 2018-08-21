package client

import (
	"cradle/walle/common"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strconv"
	"time"
)

const (
	getQyTokenUrl     string = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid="
	getQyUsersUrl     string = "https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token="
	getGitlabUsersUrl string = "https://gitlab.wallstcn.com/api/v4/users"
)

var qyToken common.QyToken
var qyUser common.QyUser
var qyUserItem common.QyUserItem

var GitlabEmailMap = make(map[string]common.GitlabUser)

var QyEmailMap = make(map[string]common.QyUserItem)

/* get token and Cache to redis*/
func GetQyToken() (string, error) {
	corpid, err := GetRedisClient().Get("corpid").Result()
	corpsecret, err := GetRedisClient().Get("corpsecret").Result()
	val, err := GetRedisClient().Get("qytoken").Result()
	/* save data into redis */
	if err != nil || len(val) == 0 {

		request := gorequest.New()
		resp, body, errs := request.Get(getQyTokenUrl + corpid + "&corpsecret=" + corpsecret).End()
		if resp.StatusCode != 200 || len(errs) != 0 {
			newError := errors.New("getQyToken resp is not 200 or errs is not null")
			return "", newError
		}

		json.Unmarshal([]byte(body), &qyToken)

		if qyToken.ErrMsg != "ok" {
			newError := errors.New("getQyToken resp is 200 but errMsg is not ok")
			fmt.Println(qyToken)
			return "", newError
		}

		err := GetRedisClient().Set("qytoken", qyToken.AccessToken, 7100*time.Second).Err()
		if err != nil {
			newError := errors.New("set redis error:" + err.Error())
			return "", newError
		}

		return qyToken.AccessToken, nil
	}

	return val, nil
}

/* get qiye weixin Users*/
func GetQyUsers() {
	token, e := GetQyToken()
	if e != nil {
		panic(e)
	}
	request := gorequest.New()
	resp, body, errs := request.Get(getQyUsersUrl + token + "&department_id=1&fetch_child=1").End()
	if resp.StatusCode != 200 || len(errs) != 0 {
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
	private_token, err := GetRedisClient().Get("private_token").Result()
	if err != nil {
		panic(err)
	}
	request := gorequest.New()
	gitlabUser := make([]common.GitlabUser, 10)
	for page := 1; page < 7; page++ {
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
	private_token, err := GetRedisClient().Get("private_token").Result()
	if err != nil {
		panic(err)
	}
	request := gorequest.New()
	resp, body, errs := request.Post(getGitlabUsersUrl + "/" + strconv.FormatInt(userId, 10) + "/block" + "?private_token=" + private_token).End()
	if resp.StatusCode != 200 && errs != nil {
		panic(errs)
	}
	return body
}
