// Go suporta _embedding_ ou _incorporação_ de structs e
// interfaces para expressar uma _composição_ de tipos
// que faça mais sentido.
// Não deve ser confundido com `//go:embed` que é uma diretiva de
// Go, introduzida na versão 1.16+ para incorporar arquivos e
// diretórios no binário de uma aplicação.

package main

import "fmt"

type base struct {
	num int
}

func (b base) descrever() string {
	return fmt.Sprintf("base com numero=%v", b.num)
}

// Uma struct `container` _incorpora_ a struct `base`.
// Uma incorporação é como um campo da struct, mas sem um nome.
type container struct {
	base
	str string
}

func main() {

	// Ao criar structs com literais, é necessário inicializar
	// a incorporação explicitamente; aqui, o tipo incorporado
	// serve como nome do campo.
	co := container{
		base: base{
			num: 1,
		},
		str: "algum nome",
	}

	// É possível acessar o campo da struct `base` diretamente
	// na variável `co`, por exemplo com `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Alternativamente, pode ser acessado com o caminho completo
	// usando o nome do tipo incorporado.
	fmt.Println("outro número:", co.base.num)

	// Como `container` incorpora `base`, os métodos de
	// `base` também se tornam métodos do `container`.
	// Aqui é invocado o método de `base` diretamente em `co`.
	fmt.Println("descreva:", co.descrever())

	type descritor interface {
		descrever() string
	}

	// Structs incorporadas com métodos podem ser usadas para
	// conceder implementações de interfaces em outras structs.
	// Aqui é visto que um `container` agora implementa a
	// interface `descritor` porque incorpora `base`.
	var d descritor = co
	fmt.Println("descritor:", d.descrever())
}
