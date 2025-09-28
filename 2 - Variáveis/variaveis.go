package main

import "fmt"

func main() {
	// Declaração de variavel direta
	var variavel1 string = "Variável 1"
	// Declaração de variavel implicita (Inferencia de tipo)
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Declaração em lote
	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)
	fmt.Println(variavel3, variavel4)

	// Declaração em lote com Inferencia de tipo
	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5, variavel6)

	// Declaração de contante
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// Pode trocar valores de variaveis sem necessidade de um auxiliar
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
