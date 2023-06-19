package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

type Msg struct {
	Cmd      string   `json:"cmd"`
	Hostname string   `json:"hostname"`
	Data     []string `json:"data"`
}

var (
	hostname  string
	chRemotes chan []string
	cont      int
	mynum     int
	nextnum   int
	next      string
	imfirst   bool
	ready     chan bool
	yach      chan bool
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
	switch msg.Cmd {
	case "connect":
		enc := json.NewEncoder(cn)
		remotes := <-chRemotes
		enc.Encode(Msg{"welcome", hostname, remotes})
		for _, remote := range remotes {
			send(remote, Msg{"update", hostname, []string{msg.Hostname}})
		}
		remotes = append(remotes, msg.Hostname)
		chRemotes <- remotes
		// fmt.Println(remotes)
	case "update":
		chRemotes <- append(<-chRemotes, msg.Data[0])
		// fmt.Println(remotes)
	case "agrawalla":
		mynum = rand.Intn(int(1e9))
		nextnum = int(1e9) + 1
		imfirst = true
		next = ""
		cont = 0
		remotes := <-chRemotes
		for _, remote := range remotes {
			send(remote, Msg{"num", hostname, []string{fmt.Sprintf("%d", mynum)}})
		}
		chRemotes <- remotes
		yach <- true
	case "num":
		<-yach
		newnum, _ := strconv.Atoi(msg.Data[0])
		if mynum > newnum {
			imfirst = false
		} else if newnum < nextnum {
			nextnum = newnum
			next = msg.Hostname
		}
		cont++
		remotes := <-chRemotes
		chRemotes <- remotes
		if cont == len(remotes) {
			if imfirst {
				fmt.Println("I'm first")
				go distributedCriticalSection()
			}
			ready <- true
		} else {
			yach <- true
		}
	case "start":
		go distributedCriticalSection()
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
			chRemotes <- append(<-chRemotes, msg.Data...)
			// fmt.Println(remotes)
		}
	}
}

func distributedCriticalSection() {
	<-ready
	fmt.Println("Me toca!")
	if next != "" {
		send(next, Msg{"start", hostname, []string{}})
	} else {
		fmt.Println("I was last :(")
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	yach = make(chan bool, 1)
	ready = make(chan bool)
	chRemotes = make(chan []string, 1)
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <hostname:port> [remote:port]\n\n", os.Args[0])
	} else if len(os.Args) == 2 {
		hostname = os.Args[1]
		chRemotes <- []string{}
		listen()
	} else if len(os.Args) == 3 {
		hostname = os.Args[1]
		chRemotes <- []string{os.Args[2]}
		go func() {
			send(os.Args[2], Msg{"connect", hostname, []string{}})
		}()
		listen()
	} else {
		for _, remote := range os.Args[1:] {
			send(remote, Msg{"agrawalla", "anonymous", []string{}})
		}
	}
}
