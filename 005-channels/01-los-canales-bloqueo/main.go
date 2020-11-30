

package main


import (
	"fmt"
)

func main() {
	// onBuffer channel (canal sin buffer)
	ca := make(chan int)

	ca <- 42

	fmt.Println(<-ca)
}

