package main

import "fmt"

var inteiro int     // 0
var texto string    // ""
var decimal float64 // 0
var booleano bool   // false

// Pointers, functions, interfaces, slices, channels, maps: nil.

func main() {
	fmt.Printf("%v, %T\n", inteiro, inteiro)
	fmt.Printf("%v, %T\n", texto, texto)
	fmt.Printf("%v, %T\n", decimal, decimal)
	fmt.Printf("%v, %T\n", booleano, booleano)
}
