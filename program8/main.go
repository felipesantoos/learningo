package main

import (
	"fmt"
)

func main() {
	teste1 := "Teste 1"
	fmt.Print(teste1, "Teste 2", "Teste", 3, "\n")
	fmt.Println(teste1, "Teste 2", "Teste", 3)
	fmt.Printf("%v %v %v %v\n", teste1, "Teste 2", "Teste", 3)

	str1 := fmt.Sprint("Hello", ", ", "World")
	str2 := fmt.Sprintln("Hello,", "World!")
	str3 := fmt.Sprintf("%s%s %s", "Hello", ",", "World!")

	fmt.Println(str1)
	fmt.Print(str2)
	fmt.Println(str3)

	// File prints
	// fmt.Fprint()
	// fmt.Fprintln()
	// fmt.Fprintf()
}
