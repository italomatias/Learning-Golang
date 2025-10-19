package main

import "fmt"

// RECEBE PONTEIRO DO NUMERO
func inverter(numero *int) {
	// RECEBE O VALOR DO PONTEIRO DO NUMERO
	*numero = *numero * -1
}

func main() {
	fmt.Println("Ponteiros")
	num := 20

	fmt.Println(num)
	// & PRA COLOCAR O ENDEREÇO DE MEMORIA DO VALOR DE NUM
	inverter(&num)
	fmt.Println(num)
}
