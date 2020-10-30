package main

import (
	"fmt"
	"net"
)

func send(num int, remote string) {
	if remote != "" {
		con, _ := net.Dial("tcp", remote)
		defer con.Close()
		fmt.Fprintf(con, "%d\n", num)
	}
}

func main() {
	nums := []int{3, 1, 4, 2, -1}
	for _, num := range nums {
		fmt.Printf("Sending %d\n", num)
		send(num, "localhost:8000")
	}
}
