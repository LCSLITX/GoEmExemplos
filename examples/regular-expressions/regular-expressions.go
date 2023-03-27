// Go oferece suporte nativo para [expressões regulares](https://pt.wikipedia.org/wiki/Express%C3%A3o_regular).
// Aqui estão alguns exemplos comuns de tarefas
// relacionadas a regEx em Go

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// Aqui é testado se um padrão corresponde a uma string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Acima foi usado o padrão de string diretamente, mas
	// para outras tarefas de regEx, é necessário compilar
	// uma struct otimizada `Regexp` struct.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Muitos métodos estão disponíveis nestas structs.
	// Aqui está uma verificação de correspondência,
	// como a que vimos antes.
	fmt.Println(r.MatchString("peach"))

	// Isto encontra a correspondência para o regexp.
	fmt.Println(r.FindString("peach punch"))

	// Isto também encontra a primeira correspondência mas
	// retorna os índices de início e final ao invés do texto.
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// A variante `Submatch` inclui informação sobre
	// a correspondência e subcorrespondências dentro
	// das correspondências. Por exemplo, isto retornará
	// informação para ambos `p([a-z]+)ch` e `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Similarmente esta forma retorna informação sobre os
	// índices das correspondências e subcorrespondências.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// A variante `All` destas funções aplicam para todas
	// as correspondências de determinado input, não apenas
	// a primeira.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Estas variantes `All`também estão disponíveis para
	// outras funções vistas anteriormente.
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Fornecer um inteiro não-negativo como segundo argumento
	// para estas funções limitará o número de correspondências.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Nos exemplos acima, foram passados argumentos do tipo
	// string e utilizados funções com string no nome, como
	// `MatchString`. Também é possível utilizar as versões
	// que aceitam `[]byte` como argumentos, deixando de usar
	// o `String` no nome da função.
	fmt.Println(r.Match([]byte("peach")))

	// Ao criar variáveis globais com expressões regulares,
	// é permitido usar a variação `MustCompile` do `Compile`.
	// `MustCompile` gera panico ao invés de retornar um erro,
	// o que torna a utilização de variáveis globais mais segura.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// O pacote `regexp` também pode ser usado para substituir
	// trechos de strings com outros valores.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// A variante `Func` permite transformar o texto correspondente
	// com determinada função.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
