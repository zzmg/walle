package main

import (
	"cradle/walle/common"
)

import (
	clt "cradle/walle/rpcclient"
)

func main() {
	common.LoadConfig("./conf/walle.yaml")
	common.Initalise()
	clt.StartService()
	clt.ClientSendEmail()

	select {}
}
