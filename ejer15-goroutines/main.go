package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLentooo("Steven")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
}

func miNombreLentooo(nombre string) {
	letras := strings.Split(nombre, "") // esto sirve para separar cada caracter
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
