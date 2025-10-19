package main

import "fmt"

type pessoa struct {
	nome  string
	idade int8
}

// METHODO DO STRUCT
// COMSEGUE CHAMAR OS DADOS DO USUÁRIO SEM PRECISAR DE PARAMAETROS
// () INICIAL MARCA QUE É UM METODO DESTE STRUCT
// METHODO PRECISA ESTAR VINCULADO A UM STRUCT PARA FUNCIONAR
func (p pessoa) salvar() {
	fmt.Println("Salvando Pessoa", p.nome)
}

// PONTEIRO PARA O STRUCT
// MUDA VALOR DO STRUCT
func (p *pessoa) aumentarIdade() {
	p.idade++
}

func main() {
	u := pessoa{"Italo", 30}
	u.salvar()
	u.aumentarIdade()
	fmt.Println(u.nome, u.idade)
}
