package test

import (
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/xinge"
	"time"
	"fmt"
	"context"
	"gitlab.wallstcn.com/wscnbackend/ivankastd/service"
	"gitlab.wallstcn.com/wscnbackend/ivankastd"
	"github.com/micro/go-micro"
)

func Test()  {
	svc := service.NewService(
		ivankastd.ConfigService{SvcName: "gitlab.wallstcn.com.xinge", SvcAddr: ":10087", EtcdAddrs: []string{"10.0.0.154:2379", "10.0.0.161:2379", "10.0.0.48:2379"}},
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	svc.Init()

	var pus xinge.PushApiClient
	pus=xinge.NewPushApiClient("gitlab.wallstcn.com.xinge",svc.Client())
	req:=new(xinge.EmailParms)
	req.Project="project"
	req.Receivers=[]string{"wangxia@wallstreetcn.com","zhangmengge@wallstreetcn.com"}
	req.Titile="titile"
	req.Content="content"
	ctx, _ := context.WithTimeout(context.Background(), (10 * time.Second))
	rsp,_:=pus.SendEmail(ctx,req)
	fmt.Println(rsp.Status)
}
