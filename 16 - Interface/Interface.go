package main

import "fmt"

// DEFINE INTERFACE
type forma interface {
	// METHODOS DA INTERFACE
	// DEVEM SER EXATAMENTE IGUAL A ESTE
	area() float64
}

// FUNC DA INTERFACE
func escreverArea(f forma) {
	fmt.Println("Area da forma é: ", f.area())
}

type quadrado struct {
	lado float64
}

// IMPLEMENTACAO DA INTERFACE EM UM STRUCT
func (q *quadrado) area() float64 {
	return q.lado * q.lado
}
func main() {
	fmt.Println("Interface")
	q := quadrado{10}
	escreverArea(&q)
}
