package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if cn, err := net.Dial("tcp", "localhost:8000"); err == nil {
		defer cn.Close()
		r := bufio.NewReader(cn)
		gin := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Tu mensaje? ")
			if msg, err := gin.ReadString('\n'); err == nil {
				fmt.Fprint(cn, msg)
				if msg[0] == '.' {
					break
				}
				if resp, err := r.ReadString('\n'); err == nil {
					fmt.Printf("Respuesta: %s", resp)
				} else {
					break
				}
			} else {
				break
			}
		}
	}
}
