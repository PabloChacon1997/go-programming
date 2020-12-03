package perrito

import (
	"testing"
)



func TestEdad(t *testing.T) {
	n := Edad(10)
	if n != 70 {
		t.Error("Expected", 70,  " Got", n)
	}
}
func TestEdad2(t *testing.T) {
	n := Edad(70)
	if n != 70 {
		t.Error("Expected", 70,  " Got", n)
	}
}
