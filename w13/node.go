package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"
)

type Player struct {
	Team string `json:"team"`
	Home string `json:"home"`
	From string `json:"from"`
}

type Message struct {
	Cmd        string `json:"cmd"`
	Hostname   string `json:"hostname"`
	Contestant Player `json:"player"`
}

type Info struct {
	hostname   string
	prev       string
	next       string
	challenger *Player
}

func listen(hostname string, chInfo chan Info) {
	if ln, err := net.Listen("tcp", hostname); err == nil {
		defer ln.Close()
		fmt.Println("Listening...")
		for {
			if cn, err := ln.Accept(); err == nil {
				go handle(cn, chInfo)
			}
		}
	}
}

func handle(cn net.Conn, chInfo chan Info) {
	defer cn.Close()
	fmt.Printf("Connection accepted from %s\n", cn.RemoteAddr())
	msg := &Message{}
	dec := json.NewDecoder(cn)
	if err := dec.Decode(msg); err == nil {
		fmt.Println(msg)
		switch msg.Cmd {
		case "jump":
			info := <-chInfo
			enc := json.NewEncoder(cn)
			if err := enc.Encode(Message{Cmd: "ok"}); err != nil {
				fmt.Printf("Can't encode OK REPLY\n%s\n", err)
			}
			player := msg.Contestant
			if info.challenger != nil {
				var loser Player
				if rand.Intn(100) >= 50 {
					loser = player
					player = *info.challenger
				} else {
					loser = *info.challenger
				}
				send(loser.Home, Message{Cmd: "send new", Hostname: info.hostname},
					func(cn net.Conn) {})
			}
			if info.next == "" || info.prev == "" {
				fmt.Println("Perdimos!")
				return
			}
			var remote string
			if player.From == info.prev {
				remote = info.next
			} else {
				remote = info.prev
			}
			player.From = info.hostname
			send(remote, Message{"jump", info.hostname, player}, func(cn2 net.Conn) {
				duration := time.Second
				if err := cn2.SetReadDeadline(time.Now().Add(duration)); err != nil {
					fmt.Printf("SetReadDeadline failed:\n%s\n", err)
					panic("OMG!")
				}

				dec := json.NewDecoder(cn2)
				var msg2 *Message
				if dec.Decode(msg2); err == nil {
					fmt.Println("Se supone que recibimos OK")
				} else {
					if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
						if msg.Hostname == info.next {
							info.challenger = &msg.Contestant
							chInfo <- info
						}
						fmt.Printf("read timeout:\n%s\n", err)
					} else {
						fmt.Printf("read error:\n%s\n", err)
					}
				}
			})
		case "send new":
			info := <-chInfo
			var remote string
			player := Player{Home: info.hostname, From: info.hostname}
			if info.prev == "" {
				remote = info.next
				player.Team = "Cobras"
			} else {
				remote = info.prev
				player.Team = "Leones"
			}
			send(remote, Message{"jump", info.hostname, player}, func(cn2 net.Conn) {
				duration := time.Second
				if err := cn2.SetReadDeadline(time.Now().Add(duration)); err != nil {
					fmt.Printf("SetReadDeadline failed:\n%s\n", err)
					panic("OMG!")
				}

				dec := json.NewDecoder(cn2)
				var msg2 *Message
				if dec.Decode(msg2); err == nil {
					fmt.Println("Se supone que recibimos OK")
				} else {
					if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
						if msg.Hostname == info.next {
							info.challenger = &msg.Contestant
							chInfo <- info
						}
						fmt.Printf("read timeout:\n%s\n", err)
					} else {
						fmt.Printf("read error:\n%s\n", err)
					}
				}
			})
			chInfo <- info
		}
	} else {
		fmt.Printf("Couldn't decode: %s\n", err)
	}
}

func send(remote string, msg Message, f func(cn net.Conn)) {
	if cn, err := net.Dial("tcp", remote); err == nil {
		defer cn.Close()
		enc := json.NewEncoder(cn)
		if err := enc.Encode(msg); err != nil {
			fmt.Printf("Couldn't enconde %s\n", err)
			f(cn)
		}
	} else {
		fmt.Printf("Failed to send: %s\n", err)
	}
}

func main() {
	special := flag.Bool("s", false, "The special flag for testing stuff.")
	hostname := flag.String("h", "", "IP/Hostname:port to listen on.")
	prevRemote := flag.String("p", "", "Previous node. If empty, home team 1")
	nextRemote := flag.String("n", "", "Next node. If empty, home team 2")
	flag.Parse()

	if *special { // TODO assuming that -n and -p are set, check
		go send(*nextRemote, Message{Cmd: "send new"}, func(cn net.Conn) {})
		go send(*prevRemote, Message{Cmd: "send new"}, func(cn net.Conn) {})
		time.Sleep(time.Second)
		return
	}

	if *hostname == "" || (*prevRemote == "" && *nextRemote == "") {
		flag.PrintDefaults()
		return
	}

	chRemotes := make(chan Info, 1)
	chRemotes <- Info{*hostname, *prevRemote, *nextRemote, nil}
	listen(*hostname, chRemotes)
}
