package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

type Msg struct {
	Cmd string `json:"cmd"`
	Num int `json: num`
}

func listen(hostname, remote string) {
	if ln, err := net.Listen("tcp", hostname); err == nil {
		defer ln.Close()
		for {
			if cn, err := ln.Accept(); err == nil {
				go handle(cn, remote)
			}
		}
	}
}

func handle(cn net.Conn, remote string) {
	defer cn.Close()
	dec := json.NewDecoder(cn)
	msg := &Msg{}
	dec.Decode(msg)
	fmt.Println(msg)
	if  msg.Num == 0 {
		fmt.Println("Perdí :(")
	} else {
		send(remote, msg.Num - 1)
	}
}

func send(remote string, num int) {
	if cn, err := net.Dial("tcp", remote); err == nil {
		defer cn.Close()
		enc := json.NewEncoder(cn)
		enc.Encode(Msg{"catch", num})
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <hostname:port> <remotehost:post>\n", os.Args[0])
		fmt.Printf("       or\n       %s <initialhost:port>\n", os.Args[0])
	} else if len(os.Args) == 2 {
		send(os.Args[1], rand.Intn(100));
	} else {
		listen(os.Args[1], os.Args[2])
	}
}
