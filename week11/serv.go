package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8000")
	if err == nil {
		defer ln.Close()
		fmt.Println("Server listening")
		for {
			conn, err := ln.Accept()
			if err == nil {
				fmt.Println("Connection accepted")
				go handle(conn)
			}
		}
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		str, err := r.ReadString('\n')
		if err == nil {
			fmt.Println("Recibido:", str)
			fmt.Fprint(conn, str)
		} else {
			fmt.Println("Connection closed abruptly")
			break
		}
	}
}
