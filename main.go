package main

import (
	"cradle/walle/client"
	"cradle/walle/business"
	"time"
)


func main() {
	client.LoadConfig()
	business.Bussiness()
	for {
		time.Sleep(time.Second * 10)
	}

}
