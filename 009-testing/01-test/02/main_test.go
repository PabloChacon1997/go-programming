package main

import "testing"

func TestMiSuma(t *testing.T) {
	type prueba struct {
		datos []int
		respuesta int
	}

	pruebas := []prueba {
		prueba{[]int {2,4,6}, 12},
		prueba{[]int {1,1,1}, 3},
		prueba{[]int {2,5,1}, 8},
	}

	for _, x := range pruebas {
		v := miSuma(x.datos...)
		if v != x.respuesta {
			t.Error("Expected: ", x.respuesta, "Got: ", v)
		}
	}
	
}
