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

type Vote struct {
	Vote   string
	Weight int
}

var (
	chRemotes chan []string
	hostname  string

	ready2listen chan bool
	decisions    map[string]Vote
	cont         int
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
		chRemotes <- remotes
	case "update":
		remotes := append(<-chRemotes, msg.Data...)
		chRemotes <- remotes
	case "start concensus":
		weight := rand.Intn(10) + 1
		if rand.Intn(100)%2 == 0 {
			decisions[hostname] = Vote{"attack", weight}
		} else {
			decisions[hostname] = Vote{"retreat", weight}
		}
		fmt.Println(hostname, decisions[hostname])
		cont = 0
		remotes := <-chRemotes
		for _, friend := range remotes {
			send(friend, Msg{"decision",
				hostname,
				[]string{
					decisions[hostname].Vote,
					strconv.Itoa(decisions[hostname].Weight)}})
		}
		chRemotes <- remotes
		ready2listen <- true
	case "decision":
		<-ready2listen
		weight, _ := strconv.Atoi(msg.Data[1])
		decisions[msg.Addr] = Vote{msg.Data[0], weight}
		cont++
		remotes := <-chRemotes
		chRemotes <- remotes
		if cont == len(remotes) {
			contAtack := 0
			contFallb := 0
			for _, decision := range decisions {
				if decision.Vote == "attack" {
					contAtack += decision.Weight
				} else {
					contFallb += decision.Weight
				}
			}
			if contAtack < contFallb {
				fmt.Println(hostname, "RETIRADA!!")
			} else {
				fmt.Println(hostname, "ATACAR!!!!")
			}
		} else {
			ready2listen <- true
		}
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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <hostname:port> <remotehost:post>\n", os.Args[0])
		fmt.Printf("       or\n       %s <initialhost:port>\n", os.Args[0])
	} else {
		hostname = os.Args[1]
		chRemotes = make(chan []string, 1)
		ready2listen = make(chan bool)
		decisions = make(map[string]Vote)
		if len(os.Args) > 4 {
			for _, remote := range os.Args[1:] {
				send(remote, Msg{"start concensus", "starter", []string{}})
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
