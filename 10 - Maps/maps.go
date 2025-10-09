package main

import "fmt"

func main() {
	fmt.Println("MAPs")
	// Criação de map
	// Chave e valor precisam ser do mesmo tipo
	usuario := map[string]string{
		"nome":      "Italo",
		"sobrenome": "Reis"}

	fmt.Println(usuario)
	// Acessar valor do MAP
	fmt.Println(usuario["nome"])
	// Deletar chave do map
	delete(usuario, "sobrenome")
	fmt.Println(usuario)
	// Adicionar chave
	usuario["Idade"] = "31"
	fmt.Println(usuario)

}
