package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	Cmd      string   `json:"cmd"`
	Hostname string   `json:"hostname"`
	Data     []string `json:"data"`
}

type Remote struct {
	prev, next string
}

func listen(hostname string, chRemotes chan Remote) {
	if ln, err := net.Listen("tcp", hostname); err == nil {
		defer ln.Close()
		fmt.Println("Listening...")
		for {
			if cn, err := ln.Accept(); err == nil {
				go handle(cn, chRemotes)
			}
		}
	}
}

func handle(cn net.Conn, chRemotes chan Remote) {
	defer cn.Close()
	fmt.Printf("Connection accepted from %s\n", cn.RemoteAddr())
	var msg Message
	dec := json.NewDecoder(cn)
	if err := dec.Decode(&msg); err == nil {
		fmt.Println(msg)
	} else {
		fmt.Printf("Couldn't decode: %s\n", err)
	}
}

func send(remote string, msg Message) {
	if cn, err := net.Dial("tcp", remote); err == nil {
		defer cn.Close()
		enc := json.NewEncoder(cn)
		if err := enc.Encode(msg); err != nil {
			fmt.Printf("Couldn't enconde %s\n", err)
		}
	} else {
		fmt.Printf("Failed to send: %s\n", err)
	}
}

func main() {
	special := flag.Bool("s", false, "The special flag for testing stuff.")
	hostname := flag.String("h", "", "tdb")
	prevRemote := flag.String("p", "", "prev node")
	nextRemote := flag.String("n", "", "next node")
	flag.Parse()

	if *special && *nextRemote != "" {
		send(*nextRemote, Message{Cmd: "Test", Hostname: "Tester hoster"})
		return
	}

	if *hostname == "" || (*prevRemote == "" && *nextRemote == "") {
		flag.PrintDefaults()
		return
	}

	chRemotes := make(chan Remote, 1)
	chRemotes <- Remote{*prevRemote, *nextRemote}
	listen(*hostname, chRemotes)
}
