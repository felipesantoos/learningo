package main

import (
	"fmt"
)

func main() {
	print()
	variables()
}

func print() {
	_, errors := fmt.Println("Hello, World!", "Olá, Mundo!")
	fmt.Println(errors)
}

func variables() {
	number := 1
	text := "This is a text."
	state := true

	fmt.Println(number, text, state)
}
