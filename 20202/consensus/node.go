package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// Msg lorem ipsum
type Msg struct {
	Addr   string `json:"addr"`
	Option string `json:"option"`
}

func main() {
	var option string
	var attack, retreat int
	local := os.Args[1]
	remotes := os.Args[2:]
	ch := make(chan Msg)

	go server(local, remotes, ch)

	fmt.Print("Your option:")
	fmt.Scanf("%s", &option)
	sendAll(option, local, remotes)

	if option == "attack" {
		attack = 1
	} else {
		retreat = 1
	}
	for range remotes {
		msg := <-ch
		if msg.Option == "attack" {
			attack++
		} else {
			retreat++
		}
	}
	if attack > retreat {
		fmt.Println("ATTACK!!")
	} else {
		fmt.Println("run...")
	}
}

func server(local string, remotes []string, ch chan Msg) {
	if ln, err := net.Listen("tcp", local); err == nil {
		defer ln.Close()
		fmt.Printf("Listening on %s\n", local)
		for {
			if conn, err := ln.Accept(); err == nil {
				fmt.Printf("Connection accepted from %s\n", conn.RemoteAddr())
				go handle(conn, local, remotes, ch)
			}
		}
	}
}

func handle(conn net.Conn, local string, remotes []string, ch chan Msg) {
	defer conn.Close()
	dec := json.NewDecoder(conn)
	var msg Msg
	if err := dec.Decode(&msg); err == nil {
		fmt.Printf("Message: %v\n", msg)
		ch <- msg
	}
}

func sendAll(option, local string, remotes []string) {
	for _, remote := range remotes {
		send(local, remote, option)
	}
}

func send(local, remote, option string) {
	if conn, err := net.Dial("tcp", remote); err == nil {
		enc := json.NewEncoder(conn)
		if err := enc.Encode(Msg{local, option}); err == nil {
			fmt.Printf("Sending %s to %s\n", option, remote)
		}
	}
}
