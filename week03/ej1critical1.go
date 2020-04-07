package main

import "fmt"

var turn = 1

func P() {
	for {
		fmt.Println("P NCS 1")
		fmt.Println("P NCS 2")
		fmt.Println("P NCS 3")
		for turn != 1 {
		}

		fmt.Println("P CRITICAL 1")
		fmt.Println("P CRITICAL 2")
		fmt.Println("P CRITICAL 3")

		turn = 2
	}
}
func Q() {
	for {
		fmt.Println("	Q NCS 1")
		fmt.Println("	Q NCS 2")
		fmt.Println("	Q NCS 3")
		for turn != 2 {
		}

		fmt.Println("	Q CRITICAL 1")
		fmt.Println("	Q CRITICAL 2")
		fmt.Println("	Q CRITICAL 3")

		turn = 1
	}
}

func main() {
	go P()
	Q()
}
