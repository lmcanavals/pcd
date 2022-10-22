package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	if ln, err := net.Listen("tcp", "localhost:8000"); err == nil {
		defer ln.Close()
		if cn, err := ln.Accept(); err == nil {
			defer cn.Close()
			r := bufio.NewReader(cn)
			for i := 0; i < 5; i++ {
				if msg, err := r.ReadString('\n'); err == nil {
					fmt.Print(msg)
				}
			}
		}
	}
}
