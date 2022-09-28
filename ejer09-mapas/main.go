// mapas
package main

import "fmt"

func main() {
	paises := make(map[string]string, 5) // el 5 es el espacio reservado de memoria

	paises["Mexico"] = "D.F"
	paises["Paises Bajos"] = "Amsterdam"
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Tolima":      39,
	}
	fmt.Println(campeonato)

	campeonato["River Plate"] = 25
	campeonato["Tolima"] = 40
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de : %d\n", equipo, puntaje)
	}

	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
