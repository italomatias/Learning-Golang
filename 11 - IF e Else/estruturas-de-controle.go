package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")
	num := 10

	if num > 10 {
		fmt.Println("Maior que 10")
	} else {
		fmt.Println("Menor igual á 10")
	}

	//IF INIT
	// IF COM VARIAVEL CRIADA NO IF
	// VAR IF INIT SÓ EXISTEM DENTRO DO IF
	if num2 := num; num2 > 10 {
		fmt.Println("Maior que 10")
	} else {
		fmt.Println("Menor que 10")
	}

}
