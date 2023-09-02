package main

import (
	"fmt"

	"github.com/vladfreishmidt/puppy"
)

func main() {
	fmt.Println("Hello, Gophers! 😊")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

	s1 := puppy.BigBark()
	s2 := puppy.BigBarks()

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("Working on versions")

}
