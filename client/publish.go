package client

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func FileSaveRedis() {
	//var hostKey ssh.PublicKey
	key, err := ioutil.ReadFile("/conf/id_rsa")
	if err != nil {
		fmt.Printf("unable to read private key: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		fmt.Printf("unable to parse private key: %v", err)
	}
	config := &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	if err != nil {
		fmt.Printf("unable to connect: %v", err)
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
		panic(err)
	}

	buf := bufio.NewReader(f)
	var mailKey []string

	for {
		a, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {

		}
		if  len(a) == 0{
			continue
		}
		mailKey = append(mailKey, string(a))
	}
	for i := 1; i < len(mailKey)+1; i++ {
		if i&1 == 1 {
			GetRedisClient().SetNX(mailKey[i-1], mailKey[i], 7200*time.Second)
		}
	}
	os.Remove("auth.txt")

}
