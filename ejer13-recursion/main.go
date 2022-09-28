package main

import "fmt"

func main() {
	exponencia(2)
}

// recursion es una fucnion que se llama a si misma
func exponencia(numero int) {
	if numero > 1000000 {
		return
	}
	fmt.Println(numero)
	exponencia(numero * 2)
}
