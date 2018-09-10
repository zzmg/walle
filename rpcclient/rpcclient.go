package rpcclient

import (
	"context"
	"fmt"
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/xinge"
	"time"
	"github.com/micro/go-micro"
	ivksvc "gitlab.wallstcn.com/wscnbackend/ivankastd/service"
	"cradle/walle/common"

)
var pus xinge.PushApiClient

func Init(svc micro.Service) {
	pus = xinge.NewPushApiClient("gitlab.wallstcn.com.xinge", svc.Client())
}

func StartService() {
	svc := ivksvc.NewService(common.GlobalConf.Micro)
	svc.Init()
	Init(svc)
}
func ClientSendEmail() {

	req := new(xinge.EmailParms)
	req.Receivers = []string{"wangxia@wallstreetcn.com", "zhangmengge@wallstreetcn.com"}
	req.Titile = "Title"
	req.Content = "Content"
	ctx, _ := context.WithTimeout(context.Background(), (10 * time.Second))
	rsp, _ := pus.SendEmail(ctx, req)
	fmt.Println(rsp.Status)
}
