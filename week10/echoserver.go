package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(cn net.Conn, i int) {
	defer cn.Close()
	r := bufio.NewReader(cn)
	for {
		if msg, err := r.ReadString('\n'); err == nil {
			fmt.Printf("Client %d said: %s", i, msg)
			if msg[0] == '.' {
				break
			}
			fmt.Fprintf(cn, "Re: %s", msg)
		} else {
			break
		}
	}
}

func main() {
	if ln, err := net.Listen("tcp", "localhost:8000"); err == nil {
		defer ln.Close()
		i := 0
		for {
			if cn, err := ln.Accept(); err == nil {
				go handle(cn, i)
				i++
			}
		}
	}
}
