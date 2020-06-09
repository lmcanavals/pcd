package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*func sumar(a, b []float64) []float64 {
	c := make([]float64, len(a))

	for i := range a {
		c[i] = a[i] + b[i]
	}
	return c
}*/
func sumar(a, b []float64) []float64 { // SPMD
	c := make([]float64, len(a))

	numproc := 6
	end := make(chan bool)
	for i := 0; i < numproc; i++ {
		go func(id int) {
			for i := id; i < len(a); i += numproc {
				c[i] = a[i] + b[i]
			}
			end <- true
		}(i)
	}
	for i := 0; i < numproc; i++ {
		<-end
	}

	return c
}
func main() {
	a := make([]float64, 1<<20)
	b := make([]float64, 1<<20)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := range a {
		a[i] = rand.Float64() * 1000
		b[i] = rand.Float64() * 1000
	}

	c := sumar(a, b)

	fmt.Println(a[5000:5020])
	fmt.Println(b[5000:5020])
	fmt.Println(c[5000:5020])
}
