package main

import (
	"cradle/walle/client"
	"time"
	"cradle/walle/test"
)


func main() {
	client.LoadConfig()
	//business.Bussiness()
	test.Test()
	for {
		time.Sleep(time.Second * 10)
	}

}
