package main

import (
	"fmt"
	"github.com/pablo/go-programming/008-NinjaNivel12/01/perro"
)

type canino struct {
	nombre string
	edad int
}


func main() {
	c1 := canino {
		nombre: "Cheester",
		edad: perro.Edad(2),
	}

	fmt.Println(c1)


}