package main

import (
	"fmt"
	"sync"
)

var roomEmpty sync.Mutex
var mutex sync.Mutex

var readers int

func Writer() {
	for {
		roomEmpty.Lock()
		fmt.Println("Writing writing writing (1)")
		fmt.Println("Writing writing writing (2)")
		fmt.Println("Writing writing writing (3)")
		roomEmpty.Unlock()
	}
}

func Reader(name string) {
	for {
		mutex.Lock()
		readers++
		if readers == 1 {
			roomEmpty.Lock()
		}
		mutex.Unlock()

		fmt.Println(name, "doing doing doing (1)")
		fmt.Println(name, "doing doing doing (2)")
		fmt.Println(name, "doing doing doing (3)")

		mutex.Lock()
		readers--
		if readers == 0 {
			roomEmpty.Unlock()
		}
		mutex.Unlock()
	}
}

func main() {
	go Reader("  Willington")
	go Reader("    Diego")
	go Reader("      Joaquin")
	go Reader("        Jorge")
	Writer()
}
