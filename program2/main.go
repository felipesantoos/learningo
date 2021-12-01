package main

import (
	"fmt"
)

func main() {
	dez := 10
	texto := "\"Isto é um texto.\""

	fmt.Printf("%v é do tipo %T\n", dez, dez)
	fmt.Printf("%v é do tipo %T\n", texto, texto)
}
