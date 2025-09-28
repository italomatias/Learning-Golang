package main

import "fmt"

func main() {
	fmt.Println(somar(10, 20))

	var f = func() {
		fmt.Println("Função F")
	}

	f()

	// Consumindo os dois resultados de uma função
	soma, sub := calculos(10, 20)
	fmt.Println(soma, sub)

	// Usa apenas 1 retorno da função
	// _ permite que consuma a função ignorando o segundo retorno
	soma2, _ := calculos(10, 20)
	fmt.Println(soma2)

}

func somar(n1 int, n2 int) int {
	return n1 + n2
}

// Parametros do mesmo tipo podem ser assim
// Go aceita mais de um retorno por função
func calculos(n1, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}
