package main

import (
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/user"
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/delegate"
	"gitlab.wallstcn.com/wscnbackend/ivankastd/service"
	"gitlab.wallstcn.com/wscnbackend/ivankastd"
	"github.com/micro/go-micro"
	"time"
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/xinge"
	"context"
	"cradle/walle/client"
	"cradle/walle/common"
	"fmt"
	"strings"
)

var (
	UIC            pbuser.InternalClient
	UINFO          pbuser.UserClient
	ShortUriClient delegate.ShortUriClient
)

var Push xinge.PushApiClient

func StartClient() {
	svc := service.NewService(
		ivankastd.ConfigService{SvcName: "gitlab.wallstcn.com.walle", SvcAddr: ":10087", EtcdAddrs: []string{"10.0.0.154:2379", "10.0.0.161:2379", "10.0.0.48:2379"}},
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	svc.Init()
	//UIC = pbuser.NewInternalClient(std.UserSvcName, svc.Client())
	//UINFO = pbuser.NewUserClient(std.UserSvcName, svc.Client())
	//ShortUriClient = delegate.NewShortUriClient(std.DelegateSvcName, svc.Client())
	Push = xinge.NewPushApiClient("gitlab.wallstcn.com.xinge", svc.Client())

}

func main() {
	client.GetQyUsers()
	client.GetGitlabUsers()
	client.FileSaveRedis()

	leaveUser := make(map[string]common.GitlabUser)
	var leaveUserList []string

	//QyUsers and GitlabUsers
	for key, val := range client.GitlabEmailMap {
		if _, ok := client.QyEmailMap[key]; !ok && client.GitlabEmailMap[key].External == false && strings.Contains(key, "wallstreetcn.com") && client.GitlabEmailMap[key].State == "active" && !strings.Contains(client.GitlabEmailMap[key].Name, "junzhi") && client.GitlabEmailMap[key].Name != "wallstreetcn" {
			fmt.Println("Users who need to be blocked on gitlab: " + val.Name + "  " + val.Email)
			//println(client.BlockGitlabUsers(val.Id))
			leaveUser[key] = val
			leaveUserList = append(leaveUserList, key)

		}
	}

	//QyUsers and sshkey
	redisList := client.GetRedisClient().Keys("*wall*").Val()
	var leaverUserPublish []string
	for _, val := range redisList {
		if _, ok := client.QyEmailMap[val[1:]]; !ok {
			fmt.Println("Users who need to be deleted on publish machine: " + val[1:])
			leaverUserPublish = append(leaverUserPublish, val[1:])
		}
	}
	StartClient()

	emailList := []string{"zhangmengge@wallstreetcn.com"}
	emailParams := xinge.EmailParms{}
	for _, val := range leaveUserList {
		emailParams.Content = "Users who need to be deleted on gitlab: " + val
	}
	for _, val := range leaverUserPublish {
		emailParams.Content = "Users who need to be deleted on publish machine: " + val
	}
	emailParams.Titile = "Users who need to be deleted"
	emailParams.Receivers = emailList
	emailParams.Project = "delete me"

	status, err := Push.SendEmail(context.Background(), &emailParams)
	if err != nil {
		fmt.Println("error in email-sending: ", err.Error())
	}
	fmt.Println("email-sending status: ", status.Status)

	for {
		time.Sleep(time.Second * 10)
	}

}
