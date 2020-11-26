
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Numero de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de GORUTINAS: %v\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hola desde la primera go rutina.")
		wg.Done()
	}()

	fmt.Printf("Numero de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de GORUTINAS: %v\n", runtime.NumGoroutine())

	go func() {
		fmt.Println("Hola desde la segunda go rutina.")
		wg.Done()
	}()

	fmt.Printf("Numero de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de GORUTINAS: %v\n", runtime.NumGoroutine())

	wg.Done()

	fmt.Println("Apunto de finalizar.")
	fmt.Printf("Numero de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de GORUTINAS: %v\n", runtime.NumGoroutine())
}


