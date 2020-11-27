
package main


import (
	"fmt"
	"runtime"
)


func main() {
	fmt.Printf("El OS: %v\nEl ARCH: %v\n", runtime.GOOS, runtime.GOARCH)
}

