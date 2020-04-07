package main

import "fmt"

var wantp = false
var wantq = false

func P() {
	for {
		fmt.Println("P NCS 1")
		fmt.Println("P NCS 2")
		fmt.Println("P NCS 3")
		wantp = true
		for wantq {
		}

		fmt.Println("P CRITICAL 1")
		fmt.Println("P CRITICAL 2")
		fmt.Println("P CRITICAL 3")

		wantp = false
	}
}

func Q() {
	for {
		fmt.Println("	Q NCS 1")
		fmt.Println("	Q NCS 2")
		fmt.Println("	Q NCS 3")
		wantq = true
		for wantp {
		}

		fmt.Println("	Q CRITICAL 1")
		fmt.Println("	Q CRITICAL 2")
		fmt.Println("	Q CRITICAL 3")

		wantq = false
	}
}

func main() {
	go P()
	Q()
}
