package main

import "fmt"

// EXECUTA FUNC ANONIMA
// CHAMAR O () NO FINAL DA DECLARACAO
func main() {
	func() {
		fmt.Println("FUNC ANONIMA")
	}()
}
