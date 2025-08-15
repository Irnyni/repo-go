package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	b := []int{1, 2, 3, 4, 5, 6}
	a := soma(b)
	esperado := 21
	if a != esperado {
		t.Error("Expected:", a, "Got", esperado)
	}
}
