package main

import "fmt"

type errPer struct {
	info string
}

func (ep errPer) Error() string {
	return fmt.Sprintf("a qui esta el error: %v", ep.info)
}

func main() {
	e1 := errPer {
		info: "Necesito dormir mas",
	}
	foo(e1)
	
}

func foo(e error) {
	fmt.Println("foo corrio - ", e, "\n")
}