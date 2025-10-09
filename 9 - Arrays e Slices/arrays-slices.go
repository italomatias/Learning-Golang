package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string // Todos os itens do array devem ser do mesmo tipo
	array1[0] = "Italo"
	fmt.Println(array1)

	array2 := [5]string{"Pos1", "Pos2", "Pos3", "Pos4", "Pos5"}
	fmt.Println(array2)

	// Define array com a qtd de itens inseridos
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	fmt.Println("SLICE")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Adicionando um item ao slice
	// append adiciona o item e retorna um novo slice
	slice = append(slice, 6)
	fmt.Println(slice)

	// Criar Slice a partir de 1 array
	// Cria o Slice da pos 1 a 3 do array
	// O ultimo valor é o valor de parada do slice
	// Não copia o valor, realiza um ponteiro
	slice2 := array2[1:4]
	fmt.Println(slice2)
	array2[1] = "POS ALTERADA"
	fmt.Println(slice2)

}
