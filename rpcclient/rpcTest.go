package rpcclient

import (
	"context"
	"fmt"
	"gitlab.wallstcn.com/wscnbackend/ivankaprotocol/xinge"
	"time"
)

func ClientSendEmail() {
	req := new(xinge.EmailParms)
	req.Project = "Project"
	req.Receivers = []string{"wangxia@wallstreetcn.com", "zhangmengge@wallstreetcn.com"}
	req.Titile = "Titile"
	req.Content = "Content"
	ctx, _ := context.WithTimeout(context.Background(), (10 * time.Second))
	rsp, _ := pus.SendEmail(ctx, req)
	fmt.Println(rsp.Status)
}
