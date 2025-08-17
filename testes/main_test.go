package main

import (
	"testing"
)

type Teste struct {
	Data []int
	Resp int
}

func TestSoma(t *testing.T) {
	b := []int{1, 2, 3, 4, 5, 6}
	a := soma(b)
	esperado := 21
	if a != esperado {
		t.Error("Expected:", a, "Got", esperado)
	}
}
func TestTabela(t *testing.T) {
	testes := []Teste{
		Teste{Data: []int{2, 2, 2}, Resp: 6},
		Teste{Data: []int{3, 3, 3}, Resp: 10},
		Teste{Data: []int{4, 4, 4}, Resp: 12},
	}
	for _, v := range testes {
		somatoria := soma2(v.Data...)
		if somatoria != v.Resp {
			t.Error("Expected:", v.Resp, "Got:", somatoria)
		}

	}

}
