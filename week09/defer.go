package main

import "fmt"

func main() {
	defer fmt.Println("SI ésto sale al final, todos jalan")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i, "hola")
	}

	defer fmt.Println("Que tal")
	fmt.Println("No esperabas que ésto salga primero")
}
