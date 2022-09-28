package main

import "fmt"

func main() {
	// while
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// for normal
	for i = 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// for poco comun :V
	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var j = 0
	for j < 10 {
		fmt.Printf("\n Valor de j: %d", j)
		if j == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			j = j * 2
			continue
		}
		fmt.Printf(" Paso por aqui\n")
		j++
	}

	var x int = 0

RUTINA: // es como una seccion
	for x < 10 {
		if x == 4 {
			x = x + 2
			fmt.Println("Voy a Rutina sumando 2 a x")
			goto RUTINA // es como un tipo de continue
		}
		fmt.Printf("Valor de x: %d\n", x)
		x++
	}
}
