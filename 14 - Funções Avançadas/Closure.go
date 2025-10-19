package main

import "fmt"

func closure() func() {
	text := "Dentro de closure"
	funcao := func() {
		fmt.Println(text)
	}
	return funcao
}

func main() {
	text := "Dentro de main"
	fmt.Println(text)

	novaFuncao := closure()
	novaFuncao()
}
