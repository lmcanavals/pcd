package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var min int

func server(hostname, remote string, end chan bool) {
	ln, _ := net.Listen("tcp", hostname)
	defer ln.Close()
	fmt.Println("Listening!")
	for {
		con, _ := ln.Accept()
		handle(con, remote, end)
	}
}

func handle(con net.Conn, remote string, end chan bool) {
	defer con.Close()
	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(msg))
	fmt.Printf("%d from %s\n", num, con.RemoteAddr())
	if min == -1 {
		min = num
	} else if num == -1 {
		send(num, remote)
		fmt.Printf("%s -> %d\n", con.LocalAddr(), min)
		end <- true
	} else if num < min {
		send(min, remote)
		min = num
	} else {
		send(num, remote)
	}
}

func send(num int, remote string) {
	if remote != "0" {
		con, _ := net.Dial("tcp", remote)
		defer con.Close()
		fmt.Fprintf(con, "%d\n", num)
	}
}

func main() {
	var hostname, remote string
	fmt.Print("Hostname: ")
	fmt.Scanf("%s", &hostname)
	fmt.Print("Remote hostname: ")
	fmt.Scanf("%s", &remote)

	remote = fmt.Sprintf("localhost:800%s", remote)

	if hostname == "" {
		fmt.Println("I'm the starter node")
		var num int
		for {
			fmt.Print("Number to start: ")
			fmt.Scanf("%d", &num)
			send(num, remote)
			if num < 0 {
				break
			}
		}
	} else {
		hostname = fmt.Sprintf("localhost:800%s", hostname)
		end := make(chan bool)
		min = -1

		go server(hostname, remote, end)
		<-end
		fmt.Println("That's  all, folks!")
	}
}
