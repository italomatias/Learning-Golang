package main

import "fmt"

// CHAMA A FUNC INIT ANTES DA MAIN
// É UMA POR ARQUIVO
func init() {
	fmt.Println("EXECUTANDO FUNC INIT")

}
func main() {
	fmt.Println("Inicializando")
}
