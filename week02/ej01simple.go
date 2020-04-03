package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n int

func rndSleep() {
	t := time.Duration(rand.Intn(10) + 5)
	time.Sleep(t * time.Millisecond)
}

func P() {
	k1 := 1
	rndSleep()
	n = k1
}

func Q() {
	k2 := 2
	rndSleep()
	n = k2
}

func main() {
	go P()
	go Q()
	time.Sleep(20 * time.Millisecond)
	fmt.Println(n)
}
