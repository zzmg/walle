package client

import (
	"golang.org/x/crypto/ssh"
	"bytes"
	"fmt"
	"net"
	"os"
	"bufio"
	"io"
	"time"
)

func FileSaveRedis(){
	config := &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.Password("253Huaerjie!"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", "123.206.184.237:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	session, err := client.NewSession()

	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("cat /home/ubuntu/.ssh/authorized_keys"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	//fmt.Printf(b.String())

	file, _ := os.Create("auth.txt")
	b.WriteTo(file)
	f, err := os.Open("auth.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	//clientRedis := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})

	buf := bufio.NewReader(f)
	var mailKey []string

	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		mailKey = append(mailKey, string(a))
	}
	for i := 1; i < len(mailKey)+1; i++ {
		if i&1 == 1 {
			GetRedisClient().Set(mailKey[i-1], mailKey[i], 23*time.Hour)
		}
	}

}
