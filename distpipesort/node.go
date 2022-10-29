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
	Num int    `json: num`
}

func listen(hostname string, chSender chan<- Msg, chLow chan int) {
	if ln, err := net.Listen("tcp", hostname); err == nil {
		defer ln.Close()
		for {
			if cn, err := ln.Accept(); err == nil {
				go handle(cn, chSender, chLow)
			}
		}
	}
}

func handle(cn net.Conn, chSender chan<- Msg, chLow chan int) {
	defer cn.Close()
	dec := json.NewDecoder(cn)
	msg := &Msg{}
	dec.Decode(msg)
	fmt.Println(msg)
	if msg.Cmd == "num" {
		lowest := <-chLow
		if lowest == -1 {
			chLow <- msg.Num
		} else if msg.Num < lowest {
			chLow <- msg.Num
			chSender <- Msg{"num", lowest}
		} else {
			chSender <- Msg{"num", msg.Num}
			chLow <- lowest
		}
	} else {
		num := <-chLow
		if num >= 0 {
			chSender <- Msg{"end", 0}
			chLow <- -1
			fmt.Println(num)
		}
	}
}

func send(remote string, msg Msg) {
	if cn, err := net.Dial("tcp", remote); err == nil {
		defer cn.Close()
		enc := json.NewEncoder(cn)
		enc.Encode(msg)
	}
}

func senderProc(remote string, chSender <-chan Msg) {
	for msg := range chSender {
		send(remote, msg)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <hostname:port> <remotehost:post>\n", os.Args[0])
		fmt.Printf("       or\n       %s <initialhost:port>\n", os.Args[0])
	} else if len(os.Args) == 2 {
		for i := 0; i < 8; i++ {
			send(os.Args[1], Msg{"num", rand.Intn(100)})
		}
		send(os.Args[1], Msg{"end", 0})
	} else {
		chSender := make(chan Msg)
		chLow := make(chan int, 1)
		chLow <- -1
		go senderProc(os.Args[2], chSender)
		listen(os.Args[1], chSender, chLow)
	}
}
