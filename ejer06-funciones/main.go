package main

import "fmt"

func main() {
	fmt.Println(uno(5))

	numero, estado := dos(2)

	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(5, 46, 100, 2, 3, 0, 12, 1))
	fmt.Println(Calculo(5, 46, 2))
	fmt.Println(Calculo(5))
	fmt.Println(Calculo(5, 10, 20, 30, 40, 506, 70))
}

func uno(numero int) int {
	return numero * 2
}

//ejemplo tonto
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

// parametros variables
func Calculo(numero ...int) int { // ... para indicar cuando uno no sabe cuantos parametros va a recibir
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d: ", item)
	}
	return total
}
