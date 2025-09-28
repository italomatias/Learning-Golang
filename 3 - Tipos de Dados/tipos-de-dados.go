package main

import (
	"errors"
	"fmt"
)

func main() {
	// Declaração de INT (INt tem desde 8 a 64 )
	var numero int16 = 100
	fmt.Println(numero)

	// Declaração for inferencia, (Int sem definir
	//tamanho pega por padrão o do sistema operacional 32 ou 34 bits)
	numero2 := -100000000000000000
	fmt.Println(numero2)

	// Integer sem sinal ( ACEITA APENAS NUMEROS POSITIVOS )
	var numero3 uint = 100
	fmt.Println(numero3)

	// Int32 = Rune
	var numero4 rune = -123456
	fmt.Println(numero4)

	// uInt8 = Byte
	var numero5 byte = 123
	fmt.Println(numero5)

	// FLOAT32 e FLOAT64
	var numero6 float32 = 123.456
	fmt.Println(numero6)
	var numero7 float64 = 123.456
	fmt.Println(numero7)

	// STRING (Não existe Char no GOLANG)
	var str string = "texto"
	fmt.Println(str)

	// Char
	char := 'A'
	fmt.Println(char) // Transforma em INT e é a posição do mesmo na tabela ASCII

	// VALOR ZERO (VALORES INICIALIZADOS NAS VARIAVEIS QUANDO NADA É ATRIBUIDO)
	var texto string
	fmt.Println(texto) //VALOR ZERO = VAZIO

	var valor int
	fmt.Println(valor) // VALOR ZERO = 0

	// Bolleano
	var booleano bool
	fmt.Println(booleano) // VALOR ZERO = FALSE

	// TIPO DE DADO ERRO
	var erro error
	fmt.Println(erro) // VALOR ZERO nil

	// Erros no GO é diferente
	// error é o tipo da variavel
	// erros é o pacote que trata os erros no GOLANG
	var erro2 error = errors.New("Erro Interno")
	fmt.Println(erro2)
}
