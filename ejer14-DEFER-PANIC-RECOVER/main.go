package main

import "log"

func main() {
	// archivo := "prueba.txt"
	// f, err := os.Open(archivo)

	// defer f.Close()

	// if err != nil {
	// 	fmt.Println("Error abriendo el archivo")
	// 	os.Exit(1)
	// }

	ejemploPanic()
}

func ejemploPanic() {
	defer func() { // se usa una funcion anonma con defer para ejecutar varias intrucciones
		reco := recover() // recoverme trae el resultado del panic
		if reco != nil {
			log.Fatalf("ocurrio un error que genero Panic\n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor de 1")
	}
}
