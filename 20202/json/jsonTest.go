package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Pelicula is a struct
type Pelicula struct {
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
	Synopsis string `json:"synopsis"`
}

func main() {
	fmt.Println("Json examples!")
	peliculas := []Pelicula{
		{"Wonder Woman", "Patty Jenkins", 2017, "BAM!"},
		{"Joker", "Todd Phillips", 2019, "hahaha"}}

	bytes, _ := json.MarshalIndent(peliculas, "", "\t")
	fmt.Printf("%s %T\n", string(bytes), bytes)

	var pelis2 []Pelicula
	_ = json.Unmarshal(bytes, &pelis2)
	fmt.Printf("%v %T\n", pelis2, pelis2)

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(peliculas)

	dec := json.NewDecoder(os.Stdin)
	dec.Decode(&pelis2)

	fmt.Printf("%v %T\n", pelis2, pelis2)
}
