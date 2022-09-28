// Funciones anonimas y Closures
package main

import "fmt"

var Calculo func(int, int) int = func(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Suma de 5 + 7 = %d\n", Calculo(5, 7))

	// Restamos
	Calculo = func(num1, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Restamos 6 - 4 = %d\n", Calculo(6, 4))

	// Restamos
	Calculo = func(num1, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("Dividimos 10 / 2 = %d\n", Calculo(10, 2))

	Operaciones()

	// closures
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}

	fmt.Println(resultado())
}

// Closures = maneja temas de seguridad en el codigo pa que no todo mundo tenga acceso

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
