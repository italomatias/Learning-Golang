package main

// RECEBE N NUMEROS DE UMA VEZ NUMEROS ILIMITADOS
// TEM QUE SEMPRE SER O ULTIMO
// SO PODE TER 1 POR FUNC
func somar(numeros ...int) int {
	var soma int = 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func main() {
	// N NUMERO
	somar(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// ZERO NUMEROS
	somar()
}
