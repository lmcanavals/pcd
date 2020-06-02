package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pipelinesort(id int, in, out chan int) {
	min := <-in
	for num := range in {
		if num < min {
			out <- min
			min = num
		} else {
			out <- num
		}
	}
	fmt.Printf("Proceso %d: %d\n", id, min)
	close(out)
}

func main() {
	n := 10
	ch := make([]chan int, n+1)
	ch[0] = make(chan int)

	for i := 0; i < n; i++ {
		ch[i+1] = make(chan int)
		go pipelinesort(i, ch[i], ch[i+1])
	}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10; i++ {
		ch[0] <- rand.Intn(100)
	}
	close(ch[0])
	for range ch[n] {
	}
}
