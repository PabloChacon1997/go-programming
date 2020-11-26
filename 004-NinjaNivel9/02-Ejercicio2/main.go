package main

import (
	"fmt"
)

type persona struct {
	nombre string
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func (p *persona) hablar() {
	fmt.Println("Hola soy: ", p.nombre)
}

func main() {
	p1 := persona {
		nombre: "Andres",
	}
	// diAlgo(&p1)
	p1.hablar()


}

