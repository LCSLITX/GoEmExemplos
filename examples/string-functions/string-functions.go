// O pacote `strings` da biblioteca padrão de Go, fornece muitas
// funções úteis relacionadas a strings. Aqui estão alguns exemplos
// para oferecer um melhor conhecimento do pacote.

package main

import (
	"fmt"
	s "strings"
)

// Aqui é criado um apelido para `fmt.Println` pois será
// bastante utilizado abaixo.
var p = fmt.Println

func main() {

	// Aqui está uma amostra de funções disponíveis em
	// `strings`. Como estas são funções do pacote, e não
	// métodos do próprio objeto string, é necessário passar
	// a string como primeiro argumento da função.
	// Você pode descobrir mais sobre as funções de
	// [`strings`](https://pkg.go.dev/strings) na documentação
	// do pacote
	p("Contem:         ", s.Contains("test", "es"))
	p("Count:          ", s.Count("test", "t"))
	p("TemPrefixo:     ", s.HasPrefix("test", "te"))
	p("TemSufixo:      ", s.HasSuffix("test", "st"))
	p("Índice:         ", s.Index("test", "e"))
	p("Join:           ", s.Join([]string{"a", "b"}, "-"))
	p("Repete:         ", s.Repeat("a", 5))
	p("Substitui:      ", s.Replace("foo", "o", "0", -1))
	p("Substitui:      ", s.Replace("foo", "o", "0", 1))
	p("Separa:         ", s.Split("a-b-c-d-e", "-"))
	p("letraMinuscula: ", s.ToLower("TEST"))
	p("letraMaiuscula: ", s.ToUpper("test"))
}
