package main

import "fmt"

// Definindo um struct
type usuario struct {
	nome  string
	idade int8
	end   endereco
}

type endereco struct {
	rua    string
	numero int8
}

func main() {
	fmt.Println("Structs")
	// Criando um struct ele começa com valores zero caso não instanciados
	var usuario1 usuario
	fmt.Println("user com valor zero ", usuario1)
	// Atribuindo valores ao struc*t
	usuario1.nome = "Italo"
	usuario1.idade = 31
	fmt.Println(usuario1)

	endereco := endereco{"Rua A", 100}

	// SEGUNDA FORMA DE INSTANCIAR STRUCT
	usuario2 := usuario{"Elizabeth", 47, endereco}
	fmt.Println(usuario2)

	// Criando usuário sem todos os dados
	usuario3 := usuario{idade: 21}
	fmt.Println(usuario3)

}
