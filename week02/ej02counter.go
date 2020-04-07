package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rndSleep() {
	t := time.Duration(rand.Intn(10) + 5)
	time.Sleep(t * time.Millisecond)
}

var n int

func P() {
	var temp int
	for i := 0; i < 10; i++ {
		rndSleep()
		temp = n
		rndSleep()
		n = temp + 1
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	go P()
	go P()
	time.Sleep(400 * time.Millisecond)
	fmt.Println(n)
}
