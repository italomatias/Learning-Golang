package main

import "fmt"

func diaDaSemana(numero int) string {
	// NO GO AO ENTRAR NUMA CONDICAO ELE JA SAI NÃO TEM BREAK
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}

}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	default:
		return "Outro dia"
	}
}

func main() {
	fmt.Println("Switch")
	dia := diaDaSemana(200)
	dia1 := diaDaSemana2(200)
	fmt.Println(dia)
	fmt.Println(dia1)
}
