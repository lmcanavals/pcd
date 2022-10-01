package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func randSleep() {
	dur := time.Duration(rand.Intn(150) + 100)
	time.Sleep(dur * time.Microsecond)
}

type Counters struct {
	reindeers, elves int
}

func reset() {
	go func() {
		for i := 0; i < 9; i++ {
			backVacation<-true
		}
	}()
}

var wakeUp, reindeerSem, elfSem, backVacation chan bool

func Santa(ch chan Counters) {
	for {
		randSleep(); fmt.Println("Santa is sleeping like a submarine.")
		<-wakeUp
		num := <-ch
		if num.reindeers == 9 {
			// prepare Sleigh
			randSleep(); fmt.Println("Santa starts preparing the Sleigh.")
			reindeerSem<- true
		} else if num.elves == 3 {
			// help elves
			randSleep(); fmt.Println("Santa is helping elves.")
			elfSem<- true
		}
		ch<- num
	}
}

func Reindeer(name string, ch chan Counters) {
	for {
		randSleep(); fmt.Printf("%s living la vida loca in the caribean\n", name)
		<-backVacation
		num := <-ch
		num.reindeers++
		if num.reindeers == 9 {
			wakeUp<- true
		}
		ch<- num

		<-reindeerSem
		// get hitched
		randSleep(); fmt.Printf("%s getting hitched\n", name)

		num = <-ch
		num.reindeers--
		if num.reindeers == 0 {
			randSleep(); fmt.Printf("%s and everyone delivering presents\n", name)
			reset()
		} else {
			reindeerSem<- true
		}
		ch<- num
	}
}

func main() {
	ch := make(chan Counters, 1)
	wakeUp = make(chan bool)
	reindeerSem = make(chan bool)
	elfSem = make(chan bool)
	backVacation = make(chan bool, 9)
	ch<- Counters{0, 0}
	reset()
	reindeers := strings.Split(
		"Dasher Dancer Prancer Vixen Comet Cupid Donner Blitzen Rudolph", " ")
	for _, reindeer := range reindeers {
		go Reindeer(reindeer, ch)
	}
	Santa(ch)
}
