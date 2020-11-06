package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	con, _ := net.Dial("tcp", "localhost:8000")
	defer con.Close()
	r := bufio.NewReader(con)
	gin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Mensaje: ")
		msg, _ := gin.ReadString('\n')
		fmt.Fprint(con, msg)
		if msg[:4] == "stop" {
			break
		}
		msg, _ = r.ReadString('\n')
		fmt.Printf("Respuesta: %s", msg)
	}
}
