package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func pipesort(in <-chan int, out, sorted chan<- int) {
	min := <-in
	for next := range in {
		if next < min {
			out <- min
			min = next
		} else {
			out <- next
		}
	}
	sorted <- min
	close(out)
}

func main() {
	var (
		n int
		seedok bool
	)


	if len(os.Args) > 2 {
		if num, err := strconv.Atoi(os.Args[2]); err == nil {
			rand.Seed(int64(num))
			seedok = true
		}
	}
	if !seedok {
		rand.Seed(time.Now().UnixNano())
	}
	if len(os.Args) > 1 {
		if x, err := strconv.Atoi(os.Args[1]); err != nil {
			return
		} else {
			n = x
		}
	} else {
		n = 10
	}
	ch := make([]chan int, n+1)
	sorted := make(chan int)
	a := make([]int, n)
	ch[0] = make(chan int)
	for i := 0; i < n; i++ {
		ch[i+1] = make(chan int)
		go pipesort(ch[i], ch[i+1], sorted)
	}
	go func() {
		for i := 0; i < n; i++ {
			num := rand.Intn(100)
			fmt.Println(num)
			ch[0] <- num
		}
		close(ch[0])
	}()
	for i := 0; i < n; i++ {
		a[i] = <-sorted
	}
	fmt.Println(a)
}
