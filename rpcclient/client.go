package rpcclient

import (
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/xinge"

	"github.com/micro/go-micro"
)

var pus xinge.PushApiClient

func Init(svc micro.Service) {
	pus = xinge.NewPushApiClient("gitlab.wallstcn.com.xinge", svc.Client())

}
