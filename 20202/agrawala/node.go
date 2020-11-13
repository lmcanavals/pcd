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

// Msg is so cool!
type Msg struct {
	Command  string   `json:"cmd"`
	Hostname string   `json:"hostname"`
	List     []string `json:"list"`
}

var (
	friends              []string
	local, next          string
	end, ready, yach     chan bool
	mynum, cont, nextnum int
	imfirst              bool
)

func serv() {
	ln, _ := net.Listen("tcp", local)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	dec := json.NewDecoder(conn)
	var msg Msg
	if err := dec.Decode(&msg); err == nil {
		switch msg.Command {
		case "hello":
			enc := json.NewEncoder(conn)
			if err := enc.Encode(&Msg{"hey", local, friends}); err == nil {
				for _, friend := range friends {
					fmt.Println(local, "introducing", msg.Hostname, "to", friend)
					send(friend, "meet new friend", msg.List)
				}
			}
			friends = append(friends, msg.Hostname)
			fmt.Println(local, "adding", friends)

		case "meet new friend":
			friends = append(friends, msg.List...)
			fmt.Println(local, "new friend", msg.List[0])

		case "start agrawala":
			mynum = rand.Intn(1 << 20)
			imfirst = true
			next = ""
			cont = 0
			nextnum = (1 << 20) + 1
			fmt.Println(local, "start agrawala")
			for _, friend := range friends {
				send(friend, "num", []string{strconv.Itoa(mynum)})
			}
			fmt.Println(local, "All informed")
			if len(friends) == 0 {
				go distributedCriticalSection()
				ready <- true
			}
			yach <- true

		case "num":
			<-yach
			newnum, _ := strconv.Atoi(msg.List[0])
			fmt.Println(local, cont, newnum, nextnum, next)
			if newnum < mynum {
				imfirst = false
			} else if newnum > mynum && newnum < nextnum {
				nextnum = newnum
				next = msg.Hostname
			}
			cont++
			fmt.Println(local, cont, newnum, nextnum, next)
			if cont == len(friends) {
				if imfirst {
					fmt.Println(local, "I'm first")
					go distributedCriticalSection()
				}
				ready <- true
			} else {
				yach <- true
			}

		case "you can start":
			go distributedCriticalSection()

		case "finish":
			fmt.Println("That's all folks")
			end <- true
		}
	} else {
		fmt.Printf("[ERROR]: %s\n", err.Error())
	}
}

func send(remote, cmd string, params []string) {
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	enc := json.NewEncoder(conn)
	if err := enc.Encode(&Msg{cmd, local, params}); err == nil {
		fmt.Println(local, cmd, params)
		if cmd == "hello" {
			dec := json.NewDecoder(conn)
			var resp Msg
			if err := dec.Decode(&resp); err == nil {
				friends = append(friends, resp.List...)
				fmt.Println(local, "Friends", friends)
			}
		}
	}
}

func distributedCriticalSection() {
	<-ready
	fmt.Println(local, "Aaaaaaah!! Critical section!(1)")
	fmt.Println(local, "Aaaaaaah!! Critical section!(2)")
	fmt.Println(local, "Aaaaaaah!! Critical section!(3)")
	fmt.Println(local, "Aaaaaaah!! Critical section!(4)")
	fmt.Println(local, "Aaaaaaah!! Critical section!(5)")
	if next != "" {
		send(next, "you can start", []string{})
	} else {
		for _, friend := range friends {
			send(friend, "finish", []string{})
		}
		fmt.Println(local, "I've seen thing you wouldn't believe...")
		end <- true
	}
}

func main() {
	if len(os.Args) > 3 {
		for _, remote := range os.Args[1:] {
			send(remote, "start agrawala", []string{})
		}
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
		end = make(chan bool)
		ready = make(chan bool)
		yach = make(chan bool)
		local = os.Args[1]
		go serv()
		if len(os.Args) == 3 {
			remote := os.Args[2]
			friends = append(friends, os.Args[2])
			send(remote, "hello", []string{local})
		}
		<-end
		fmt.Println(local, "time to die.")
	}
}
