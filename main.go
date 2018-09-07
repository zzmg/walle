package main

import (
	"cradle/walle/common"
	ivksvc "gitlab.wallstcn.com/wscnbackend/ivankastd/service"
)

import (
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/xinge"

	clt "cradle/walle/rpcclient"
	"github.com/micro/go-micro"
)

var pus xinge.PushApiClient

func Init(svc micro.Service) {
	pus = xinge.NewPushApiClient("gitlab.wallstcn.com.xinge", svc.Client())
}

func startService() {
	svc := ivksvc.NewService(common.GlobalConf.Micro)
	svc.Init()
	Init(svc)
	clt.Init(svc)
}
func main() {
	common.LoadConfig("./conf/walle.yaml")
	common.Initalise()
	startService()
	clt.ClientSendEmail()

	select {}
}
