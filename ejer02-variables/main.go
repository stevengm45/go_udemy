package main

import (
	"fmt"
	"strconv"
)

// las crea como variables globales
var texto string
var numero int
var status bool

func main() {

	numero1, numero2, numero3, texto := 2, 5, 67, "este es mi texto"

	numero4 := 45

	var moneda float64 = 0

	numero2 = int(moneda)
	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
}
