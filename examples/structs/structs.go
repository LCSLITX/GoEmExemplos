// As _structs_ em Go são coleções tipadas de dados.
// São úteis para agrupar informações.

package main

import "fmt"

// A struct do tipo `pessoa` tem os campos `nome` e `idade`.
type pessoa struct {
	nome  string
	idade int
}

// A função `novaPessoa` constrói uma nova
// struct de pessoa com um determinado nome.
func novaPessoa(name string) *pessoa {

	// O valor atribuído para a variável local
	// será retornado pela função.
	p := pessoa{nome: name}
	p.idade = 42
	return &p
}

func main() {

	// Esta sintaxe cria uma nova `struct`.
	fmt.Println(pessoa{"Bob", 20})

	// E possível nomear os campos ao inicializar a `struct`.
	fmt.Println(pessoa{nome: "Alice", idade: 30})

	// Os campos omitidos terão valor padrão, ou zero.
	fmt.Println(pessoa{nome: "Fred"})

	// O prefixo `&` signfica um ponteiro para a `struct`.
	fmt.Println(&pessoa{nome: "Ann", idade: 40})

	// É idiomático encapsular a criação de uma nova `struct` em uma função construtora.
	fmt.Println(novaPessoa("Jon"))

	// É possível acessar campos de determinada `struct` utilizando notação de pontos.
	s := pessoa{nome: "Sean", idade: 50}
	fmt.Println(s.nome)

	// E também utilizar notação de pontos com ponteiros de `structs`
	// ponteiros serão de-referenciados automaticamente.
	sp := &s
	fmt.Println(sp.idade)

	// Structs são mutáveis.
	sp.idade = 51
	fmt.Println(sp.idade)
}
