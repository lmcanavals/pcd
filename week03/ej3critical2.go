package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rndSleep() {
	t := time.Duration(rand.Intn(1000000) + 10)
	time.Sleep(t * time.Nanosecond)
}
func myPrint(str string) {
	rndSleep()
	fmt.Println(str)
}

var wantp = false
var wantq = false

func P() {
	for {
		myPrint("P NCS 1")
		myPrint("P NCS 2")

		for wantq {
		}
		wantp = true

		myPrint("P CRITICAL 1")
		myPrint("P CRITICAL 2")
		myPrint("P CRITICAL 3")

		wantp = false
	}
}

func Q() {
	for {
		myPrint("	Q NCS 1")
		myPrint("	Q NCS 2")

		for wantp {
		}
		wantq = true

		myPrint("	Q CRITICAL 1")
		myPrint("	Q CRITICAL 2")
		myPrint("	Q CRITICAL 3")

		wantq = false
	}
}

func main() {
	go P()
	Q()
}
