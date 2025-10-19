package main

import "fmt"

// RETORNO NOMEADO
// DEFINE O NOME DO RETORNO E NAO PRECISA DECLARAR A VAR
// E NEM PRECISA PASSAR A MESMA NO RETORNO.
func calcule(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, sub := calcule(10, 20)
	fmt.Println(soma, sub)
}
