package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Msg struct {
	Cmd string `json:"cmd"`
	Hostname string `json:"hostname"`
	Data []string `json:"data"`
}

var (
	hostname string
	remotes []string
)

func listen() {
	if ln, err := net.Listen("tcp", hostname); err == nil {
		defer ln.Close()
		for {
			if cn, err := ln.Accept(); err == nil {
				go handle(cn)
			}
		}
	}
}

func handle(cn net.Conn) {
	defer cn.Close()
	dec := json.NewDecoder(cn)
	var msg Msg
	dec.Decode(&msg)
	fmt.Println(msg)
	switch (msg.Cmd) {
	case "connect":
		enc := json.NewEncoder(cn)
		enc.Encode(Msg{"welcome", hostname, remotes})
		for _, remote := range remotes {
			send(remote, Msg{"update", hostname, []string{msg.Hostname}})
		}
		remotes = append(remotes, msg.Hostname)
		fmt.Println(remotes)
	case "update":
		remotes = append(remotes, msg.Data[0])
		fmt.Println(remotes)
	}
}

func send(remote string, msg Msg) {
	if cn, err := net.Dial("tcp", remote); err == nil {
		enc := json.NewEncoder(cn)
		enc.Encode(msg)
		if msg.Cmd == "connect" {
			dec := json.NewDecoder(cn)
			var msg Msg
			dec.Decode(&msg)
			remotes = append(msg.Data, remote)
			fmt.Println(remotes)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <hostname:port> [remote:port]\n\n", os.Args[0])
	} else if len(os.Args) == 2 {
		hostname = os.Args[1]
		listen()
	} else {
		hostname = os.Args[1]
		go func() {
			send(os.Args[2], Msg{"connect", hostname, []string{}})
		}()
		listen()
	}
}
