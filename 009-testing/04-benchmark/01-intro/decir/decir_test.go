package decir


import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {
	s := Saludar("Andres")
	if s != "Bienvenido querido: Andres" {
		t.Error("Expected","Bienvenido querido: Andres" ," Got: ",s)
	}
}


func ExampleSaludar() {
	fmt.Println(Saludar("Andres"))
}

func BenchmarckSaludar(b *testing.B) {
	for i := 0 ; i < b.N; i++ {
		Saludar("Andres")
	}
}

