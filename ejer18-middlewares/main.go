package main

import "fmt"

/*Los middlewares son interceptores que permiten ejecutar instrcucciones comunes a varias funciones
que reciben y devuelven los mismos tipos de variables */

func main() {
	fmt.Println("inicio")

	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(12, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 13)
	fmt.Println(result)

}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

// middlewares
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
