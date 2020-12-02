// package mate ayuda a saber que sabes sumar
package mate

// Suma suma un acantidad limitada de numeros
func Suma(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}

	return s
}
