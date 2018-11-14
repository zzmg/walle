package main

import (
	"cradle/walle/common"
	clt "cradle/walle/rpcclient"
	"cradle/walle/models"
)

func main() {
	common.LoadConfig("/conf/walle.yaml")
	common.Initalise()
	clt.StartService()
	clt.ClientSendEmail()
	service
	defer  models.CloseDB()
	select {}
}

func startService() {
	svc := ivksvc.NewService(common.GlobalConf.Micro)
	svc.Init()
	rpcserver.Init(svc)

}