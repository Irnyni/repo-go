package main

import (
	"fmt"
	"testing"
)

type Teste struct {
	Data []int
	Resp int
}

func Example() {
	fmt.Println(Soma2(1, 2, 3, 4, 5, 6))
	//Output:
	//21

}
func TestSoma(t *testing.T) {
	b := []int{1, 2, 3, 4, 5, 6}
	a := soma(b)
	esperado := 21
	if a != esperado {
		t.Error("Expected:", a, "Got", esperado)
	}
}
func BenchmarkNada(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{1, 2, 3, 4, 5, 6}
		b := Soma2(a...)
		fmt.Printf("%v", b)
	}
}
func TestTabela(t *testing.T) {
	testes := []Teste{
		Teste{Data: []int{2, 2, 2}, Resp: 6},
		Teste{Data: []int{3, 3, 3}, Resp: 9},
		Teste{Data: []int{45, 5, 5}, Resp: 55},
		Teste{Data: []int{4, 3, 4}, Resp: 11},
		Teste{Data: []int{4, 1, 4}, Resp: 9},
		Teste{Data: []int{4, 4, 4}, Resp: 12},
		Teste{Data: []int{4, 2, 4}, Resp: 10},
		Teste{Data: []int{41, 4, 4}, Resp: 49},
	}
	for _, v := range testes {
		somatoria := Soma2(v.Data...)
		if somatoria != v.Resp {
			t.Error("Expected:", v.Resp, "Got:", somatoria)
		}

	}

}
