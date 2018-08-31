package business

import (
	"fmt"
	"math/rand"
	"time"
	"cradle/walle/client"
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/xinge"
	"cradle/walle/common"
	"strings"
	"gitlab.wallstcn.com/wscnbackend/ivankastd/service"
	"gitlab.wallstcn.com/wscnbackend/ivankastd"
	"github.com/micro/go-micro"
	"context"
)

//var (
//	UIC            pbuser.InternalClient
//	UINFO          pbuser.UserClient
//	ShortUriClient delegate.ShortUriClient
//)

var Push xinge.PushApiClient

func StartClient() {

	svc := service.NewService(
		ivankastd.ConfigService{SvcName: "gitlab.wallstcn.com.walle", SvcAddr: ":10087", EtcdAddrs: []string{"10.0.0.154:2379", "10.0.0.161:2379", "10.0.0.48:2379"}},
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	svc.Init()
	Push = xinge.NewPushApiClient("gitlab.wallstcn.com.xinge", svc.Client())

}
func Bussiness() {
	StartClient()
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

	//SSL info
	/*init var*/
	var publicVar client.PublicVar
	var sslVar client.SslVar
	publicVar.Action = "CertGetList"
	publicVar.SecretId = client.SecretId
	publicVar.SignatureMethod = "HmacSHA256"
	publicVar.Nonce = fmt.Sprintf("%d", func() int {
		rand.Seed(time.Now().Unix())
		randNum := rand.Intn(10000000)
		return randNum
	}())
	publicVar.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	publicVar.Region = "ab-shanghai"
	sslVar.Page = "1"
	sslInfo,_ := client.GetSslInfo(publicVar, sslVar)

	//grpc server


	emailList := []string{"zhangmengge@wallstreetcn.com"}
	emailParams := xinge.EmailParms{}
	for _, val := range leaveUserList {
		emailParams.Content += "Users who need to be deleted on gitlab: " + val +"\n"
	}
	for _, val := range leaverUserPublish {
		emailParams.Content += "Users who need to be deleted on publish machine: " + val +"\n"
	}
	emailParams.Titile = "Users who need to be deleted"
	emailParams.Receivers = emailList
	emailParams.Project = ""
	emailParams.Content = "this is test"
	fmt.Printf(emailParams.Content)

	status, err := Push.SendEmail(context.Background(), &emailParams)
	if err != nil {
		fmt.Println("error in email-sending: ", err.Error())
	}
	fmt.Println("email-sending status: ", status.Status)

}
