package main

import "fmt"

type pessoa struct {
	nome  string
	idade int8
	sexo  string
}

type estudante struct {
	pessoa // Só passa o nome do Struct sem especificar o tipo
	curso  string
}

func main() {
	fmt.Println("HERANÇAS")
	p1 := pessoa{"Italo", 31, "M"}
	fmt.Println(p1)

	e1 := estudante{p1, "Golang"}
	fmt.Println(e1)
}
