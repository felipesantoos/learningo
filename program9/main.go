package main

import (
	"fmt"
)

type CPF string

var cpf CPF
var spf string

func main() {
	fmt.Printf("%T\n", cpf)
	fmt.Printf("%T\n", spf)

	// Não é possível fazer cpf = spf
}
