package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo Main.")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("abacate")
	fmt.Println(erro)
}
