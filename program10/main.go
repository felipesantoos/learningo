package main

import (
	"fmt"
)

type CPF string

var cpf CPF = "014.997.830-86"
var spf string = "983.300.350-85"

func main() {
	fmt.Println("cpf:", cpf)
	fmt.Println("spf:", spf)

	aux := cpf
	cpf = CPF(spf)
	spf = string(aux)

	fmt.Println("cpf com valor de spf:", cpf)
	fmt.Println("spf com valor de cpf:", spf)
}
