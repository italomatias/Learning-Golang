package main

import "fmt"

func aprovado(nota float64) bool {
	defer recuperar()
	if nota > 6 {
		return true
	} else if nota < 6 {
		return false
	}
	// ENTRA EM PANICO E CHAMA O DEFER
	// SE NAO FOR TRATRATO FINALIZA A EXECUCAO
	panic("A MÉDIA É EXATAMENTE 6")
}

func recuperar() {
	// RECUPERA O PANIC CASO NAO TENHA ERRO SEGUE A EXECUCAO
	fmt.Println("Recuperando")
	if r := recover(); r != nil {
		fmt.Println("Recuperado com Sucesso:", r)
	}
}

func main() {
	fmt.Println("Panic and Recover")
	aprovado(6)

}
