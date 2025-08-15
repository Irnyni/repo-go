package main

import (
	"testing"
)

func Testsoma(t *testing.T) {
	b := []int{1, 2, 3, 4, 5, 6}
	a := soma(b)
	esperado := 17
	if a != esperado {
		t.Error("Expected:", a, "Got", esperado)
	}
}
