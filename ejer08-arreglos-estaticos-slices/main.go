// Arreglos estaticos y Slices
package main

import "fmt"

//var tabla [10]int
var matriz [5][7]int

func main() {
	// tabla[0] = 1
	// tabla[5] = 15
	tabla := [10]int{5, 6, 7, 12, 3, 412, 54, 61, 7, 67}

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	fmt.Println(tabla)

	matriz[3][5] = 1
	fmt.Println(matriz)

	// Slices
	matriz2 := []int{2, 5, 4} // sino le asigno longitud se convierte en un slices
	fmt.Println(matriz2)
	variante2()
	variante3()
	variante4()
}

// slices
func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[2:4]

	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d\n", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
}
