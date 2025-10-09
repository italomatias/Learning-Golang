package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var valor int = 10
	// Ponteiro para o valor *int (Int de ponteiro) &valor ponteiro para essa var
	var valor2 *int = &valor
	// Desferenciação de ponteiros (Lê o valor do endereço de memoria do ponteiro)
	fmt.Println(valor, *valor2)
}
