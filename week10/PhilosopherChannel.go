package main

import "fmt"

func philosopher(leftFork, rightFork chan bool) {
	for {
		fmt.Println("Thinking about you.")
		<-leftFork
		<-rightFork
		fmt.Println("Eating!!")
		leftFork <- true
		rightFork <- true
	}
}

func fork(chfork chan bool) {
	for {
		chfork <- true
		<-chfork
	}
}

func main() {
	n := 5
	forks := make([]chan bool, n)
	for i := range forks {
		forks[i] = make(chan bool, 1) // buffered channel or async channel
	}
	for i := 0; i < n-1; i++ {
		go philosopher(forks[i], forks[i+1])
		go fork(forks[i])
	}
	go philosopher(forks[n-1], forks[0])
	fork(forks[n-1])
}
