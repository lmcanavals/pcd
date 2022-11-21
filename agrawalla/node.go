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
	Cmd  string   `json:"cmd"`
	Addr string   `json:"addr"`
	Data []string `json:"data"`
}

var (
	chRemotes chan []string
	hostname  string
	mynum     int
	nextnum   int
	imfirst   bool
	next      string
	cont      int
	chNext    chan bool
	chReady   chan bool
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
	msg := &Msg{}
	dec.Decode(msg)
	fmt.Println(msg)
	switch msg.Cmd {
	case "connect":
		enc := json.NewEncoder(cn)
		remotes := <-chRemotes
		enc.Encode(Msg{"welcome", hostname, remotes})
		for _, remote := range remotes {
			send(remote, Msg{"update", hostname, []string{msg.Addr}})
		}
		remotes = append(remotes, msg.Addr)
		//fmt.Println(remotes)
		chRemotes <- remotes
	case "update":
		remotes := append(<-chRemotes, msg.Data...)
		//fmt.Println(remotes)
		chRemotes <- remotes
	case "critical task":
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
		chNext <- true
	case "num":
		<-chNext
		newnum, _ := strconv.Atoi(msg.Data[0])
		if newnum < mynum {
			imfirst = false
		} else if newnum > mynum && newnum < nextnum {
			nextnum = newnum
			next = msg.Addr
		}
		cont++
		remotes := <-chRemotes
		chRemotes <- remotes
		fmt.Println(cont, next, nextnum, imfirst)
		if cont == len(remotes) {
			if imfirst {
				fmt.Println("I'm first...")
				go distributedCriticalSection()
			}
			chReady <- true
		} else {
			chNext <- true
		}
	case "you are next":
		go distributedCriticalSection()
	}
}

func send(remote string, msg Msg) {
	if cn, err := net.Dial("tcp", remote); err == nil {
		defer cn.Close()
		enc := json.NewEncoder(cn)
		enc.Encode(msg)
		if msg.Cmd == "connect" {
			dec := json.NewDecoder(cn)
			msg := &Msg{}
			dec.Decode(msg)
			remotes := append(msg.Data, <-chRemotes...)
			//fmt.Println(remotes)
			chRemotes <- remotes
		}
	}
}

func distributedCriticalSection() {
	<-chReady
	fmt.Println("Realizando tarea critica...")
	if next != "" {
		send(next, Msg{"you are next", hostname, []string{}})
	} else {
		fmt.Println("I was last :(")
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <hostname:port> <remotehost:post>\n", os.Args[0])
		fmt.Printf("       or\n       %s <initialhost:port>\n", os.Args[0])
	} else {
		hostname = os.Args[1]
		chRemotes = make(chan []string, 1)
		chNext = make(chan bool)
		chReady = make(chan bool)
		if len(os.Args) > 4 {
			for _, remote := range os.Args[1:] {
				send(remote, Msg{"critical task", "starter", []string{}})
			}
		} else if len(os.Args) == 3 {
			go func() {
				chRemotes <- []string{os.Args[2]}
				send(os.Args[2], Msg{"connect", hostname, []string{}})
			}()
		} else if len(os.Args) == 2 {
			chRemotes <- []string{}
		}
		listen()
	}
}
