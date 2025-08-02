package main

import "fmt"

func main() {

	par := make(chan int)
	impar := make(chan int)
	encerrar := make(chan string)
	go enviar(par, impar, encerrar, 500)
	receber(par, impar, encerrar)

}
func enviar(p, i chan int, e chan string, n int) {
	for f := 0; f < n; f++ {
		if f%2 == 0 {
			p <- f

		} else {
			i <- f

		}
	}
	close(p)
	close(i)
	e <- "encerrar"

}

func receber(p, i chan int, e chan string) {
	for {
		select {
		case v := <-p:
			fmt.Println("par:", v)
		case v := <-i:
			fmt.Println("impar:", v)
		case <-e:
			return
		}
	}

}
