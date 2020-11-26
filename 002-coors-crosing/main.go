package main

import (
	"fmt"
	"runtime"
)



func main() {
	fmt.Printf("Primero: %v\nSegundo: %v\n",runtime.GOOS ,runtime.GOARCH)
}
