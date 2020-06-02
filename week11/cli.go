package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err == nil {
		defer conn.Close()
		var str string
		r := bufio.NewReader(conn)
		gin := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Mensaje: ")
			str, _ = gin.ReadString('\n')
			fmt.Fprint(conn, str)
			str, _ = r.ReadString('\n')
			fmt.Print("Respuesta:", str)
		}
	}
}
