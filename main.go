package main

import (
	"cradle/autoDelUser/client"
	"strings"
	"fmt"
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/user"
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/delegate"
	"gitlab.wallstcn.com/wscnbackend/ivankastd/service"
	"github.com/micro/go-micro"
	"time"
	"gitlab.wallstcn.com/wscnbackend/ivankastd"
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/xinge"
	"golang.org/x/net/context"
)

var (
	UIC            pbuser.InternalClient
	UINFO          pbuser.UserClient
	ShortUriClient delegate.ShortUriClient
)

var Push xinge.PushApiClient

func StartClient() {
	svc := service.NewService(
		ivankastd.ConfigService{SvcName: "gitlab.wallstcn.com.autoDelUser", SvcAddr: ":10087", EtcdAddrs: []string{"10.1.0.2:2379", "10.1.0.210:2379", "10.1.0.222:2379"}},
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
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
			fmt.Println("Users who need to be deleted on gitlab: " + val.Name + "  " + val.Email)
			//println(val.Id)
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

	emailList := []string{"sre@wallstreetcn.com"}
	emailParams := xinge.EmailParms{}
	for _, val := range leaveUserList {
		emailParams.Content = "Users who need to be deleted on gitlab: " + val
	}
	for _, val := range leaverUserPublish {
		emailParams.Content = "Users who need to be deleted on publish machine: " + val
	}
	emailParams.Titile = "Users who need to be deleted"
	emailParams.Receivers = emailList
	emailParams.Project = "autoDelUser"

	Push.SendEmail(context.Background(),&emailParams)
	//fmt.Println(emailParams.Content)
	for {
		time.Sleep(time.Second * 10)
	}

}

