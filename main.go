package main

import (
	"cradle/walle/common"
)

import (
	clt "cradle/walle/rpcclient"
	"cradle/walle/rpcclient"
)

func main() {
	common.LoadConfig("./conf/walle.yaml")
	common.Initalise()
	rpcclient.StartService()
	clt.ClientSendEmail()

	select {}
}
