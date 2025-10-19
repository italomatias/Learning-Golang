package main

import "fmt"

func funcao1() {
	fmt.Println("Funcao 1")
}

func funcao2() {
	fmt.Println("Funcao 2")
}

func main() {
	fmt.Println("DEFER")
	// ADIA A EXEC DA FUNC 1
	// EXECUTA NO FINAL DA APLICAÇÃO
	defer funcao1()
	funcao2()
}
