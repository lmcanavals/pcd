package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "localhost:8000")
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go handle(con)
	}
}

func handle(con net.Conn) {
	defer con.Close()

	r := bufio.NewReader(con)
	for {
		msg, _ := r.ReadString('\n')
		fmt.Printf("%s: %s", con.RemoteAddr(), msg)
		if msg[:4] == "stop" {
			break
		}
		fmt.Fprint(con, msg)
	}
}
