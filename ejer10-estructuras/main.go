// POO GO
package main

import (
	"fmt"
	"time"

	us "newuser/user"
)

type pepe struct {
	us.Usuario
}

func main() {
	user := new(pepe)
	user.AltaUsuario(1, "Steven Gonzalez Mahecha", time.Now(), true)
	fmt.Println(user.Usuario)
}
