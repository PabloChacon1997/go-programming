

package main


import (
	"fmt"
)

func main() {
	// onBuffer channel (canal sin buffer)
	c := make(chan int)

	// enviar
	go enviar(c)

	// recibir
	recibir(c)

	fmt.Println("Finalizando")
	
}

func enviar(c chan<- int) {
	c <- 42
}
func recibir(c <-chan int) {
	fmt.Println(<-c)
}
