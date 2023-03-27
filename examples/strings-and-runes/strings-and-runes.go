// Uma string em Go é um slice de bytes, ou []byte, apenas para leitura.
// A linguagem a biblioteca padrão tratam strings especialmente - como
// containers de text codificado em [UTF-8](https://en.wikipedia.org/wiki/UTF-8).
// Em outras linguagens, strings são feitas de "caracteres".
// Em Go, o conceito de caractere é chamado de `rune`, ou runa - é
// um inteiro que representa um ponto de código Unicode.
// [Este post](https://go.dev/blog/strings) é uma boa introdução.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` é uma `string` a que foi atribuida um valor literal
	// representando a palavra "olá" em língua tailandesa.
	// As strings literais em Go são textos codificados em UTF-8.
	const s = "สวัสดี"

	// Como strings são equivalentes a `[]byte`, o próximo exemplo
	// vai produzir o tamanho em `raw bytes`, ou bytes crus, constantes
	// na string.
	fmt.Println("Len:", len(s))

	// Indexar a string produz os valores em raw byte em cada
	// índice. Este loop gera o valor hexadecimal de cada um
	// dos bytes que constituem os pontos de código em `s`.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Para contar quantas _runes_ tem numa string, é possível usar
	// o pacote `utf8`. Note que o tempo de execução da função
	// `RuneCountInString` depende do tamanho da string,
	// porque é preciso decodificar cada runa UTF-8 sequencialmente.
	// Alguns caracteres do alfabeto Tailandês são representados
	// por múltiplos pontos de código UTF-8 então o resultado
	// desta contagem pode ser um pouco surpreendente.
	fmt.Println("Conta runas:", utf8.RuneCountInString(s))

	// Um loop `range` decodifica cada `rune`
	// junto com sua posição na string.
	for idx, runeValue := range s {
		fmt.Printf("%#U começa em %d\n", runeValue, idx)
	}

	// É possível atingir o mesmo resultado da iteração anterior
	// usando a função `utf8.DecodeRuneInString`.
	fmt.Println("Usando DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U começa em %d\n", runeValue, i)
		w = width

		// Isto demonstra uma runa sendo passada para uma função.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// Valores encapsulados em aspas simples são _rune literals_.
	// É possível comparar o valor da `rune` com uma _rune literal_ diretamente.
	if r == 't' {
		fmt.Println("Encontrou um 't'")
	} else if r == 'ส' {
		fmt.Println("Encontrou uma 'saw suea'")
	}
}
