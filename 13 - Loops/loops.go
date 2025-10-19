package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	// PRIMEIRO TIPO DE FOR
	i := 0
	for i < 10 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
	}
	fmt.Println("######################")

	// SEGUNDO TIPO DE FOR
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	fmt.Println("######################")
	// FOR COM RANGE
	// PODE IGNORAR O INDICE COLOCANDO _
	nomes := []string{"Italo", "Elizabeth", "Maria"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
		time.Sleep(time.Second)
	}

	// FOR COM MAP SEMPRE EXIBE A CHAVE DEPOIS O VALOR
	fmt.Println("######################")
	usuarios := map[string]string{
		"nome":  "Italo",
		"idade": "31",
	}
	for chave, valor := range usuarios {
		fmt.Println(chave, valor)
		time.Sleep(time.Second)
	}

	//LOOP INFINITO
	for {
		fmt.Println("Loop infinito")
		time.Sleep(time.Second)
	}
}
