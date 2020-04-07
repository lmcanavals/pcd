package main

var turn = 1

func P() {
	for {
		for turn != 1 {
		}
		turn = 2
	}
}
func Q() {
	for {
		for turn != 2 {
		}
		turn = 1
	}
}

func main() {
	go P()
	Q()
}
