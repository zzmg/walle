package main

import (
	"cradle/walle/common"
	clt "cradle/walle/rpcclient"
)

func main() {
	common.LoadConfig("./conf/walle_test.yaml")
	common.Initalise()
	clt.StartService()
	clt.ClientSendEmail()
	select {}
}
