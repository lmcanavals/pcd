package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(a []float32) {
	n := len(a)
	temp := make([]float32, n)
	mid := (n - 1) / 2
	i := 0
	j := mid + 1
	for k := 0; i <= mid || j < n; k++ {
		if i > mid || j < n && a[j] < a[i] {
			temp[k] = a[j]
			j++
		} else {
			temp[k] = a[i]
			i++
		}
	}
	for k := 0; k < n; k++ {
		a[k] = temp[k]
	}
}

func mergesort(a []float32) {
	if len(a) > 1 {
		mid := (len(a) - 1) / 2
		go mergesort(a[:mid+1])
		go mergesort(a[mid+1:])
		merge(a)
	}
}

func main() {
	n := 50000
	a := make([]float32, n)
	rand.Seed(1981)
	for i := range a {
		a[i] = float32(rand.Intn(1000000))
	}
	start := time.Now()
	mergesort(a)
	fmt.Println("Time elapsed ", time.Since(start))
	sorted := true
	for i := 0; i < n-1; i++ {
		if a[i] > a[i+1] {
			sorted = false
			break
		}
	}
	if sorted {
		fmt.Println("sorted!!")
	} else {
		fmt.Println("not sorted!!")
	}
}
