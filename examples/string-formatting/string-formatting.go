// Go oferece um excelente suporte para formatação de strings
// com `printf`. Aqui estão alguns exemplos comuns de
// formatação.

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go oferece várias sintaxes para impressão de várias
	// formatações diferentes. Por exemplo, este imprime
	// uma instância de uma struct, no caso a `point`.
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// Se o valor é uma struct, a variante `%+v` inclui
	// os nomes dos campos da struct.
	fmt.Printf("struct2: %+v\n", p)

	// A variante `%#v` imprime uma representação do valor,
	// conforma a sintaxe de Go, por exemplo, o trecho do código
	// fonte que produz determinado valor.
	fmt.Printf("struct3: %#v\n", p)

	// Para imprimir o tipo do valor `%T`.
	fmt.Printf("type: %T\n", p)

	// Para formatar booleanos diretamente.
	fmt.Printf("bool: %t\n", true)

	// Existem várias opções para formatar inteiros.
	// A sintaxe `%d` é utilizada para formatação padrão
	// em base-10.
	fmt.Printf("int: %d\n", 123)

	// Aqui, é impresso uma representação em binário.
	fmt.Printf("bin: %b\n", 14)

	// Este imprime o caractere correspondente a
	// um determinado inteiro.
	fmt.Printf("char: %c\n", 33)

	// `%x` fornece encoding em hexadecimal.
	fmt.Printf("hex: %x\n", 456)

	// Também há diversas opções de formatação para `floats`.
	// Para formatação em decimal básico, `%f`.
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` e `%E` formatam o float para (versões
	// levemente diferentes de) notação científica.
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// Para formatação em string básica, `%s`.
	fmt.Printf("str1: %s\n", "\"string\"")

	// Para strings com aspas duplas, `%q`.
	fmt.Printf("str2: %q\n", "\"string\"")

	// Como visto com inteiros anteriormente, `%x` renderiza
	// a string em base-16, com dois caracteres de saídas (output)
	// por byte de entrada (input).
	fmt.Printf("str3: %x\n", "hex this")

	// Para imprimir a representação de um ponteiro, `%p`.
	fmt.Printf("pointer: %p\n", &p)

	// Ao formatar números, frequentemente será útil
	// controlar a largura e precisão da figura resultante.
	// Para especificar a largura de um inteiro, utiliza-se
	// um número depois do simbolo de porcentagem `%`.
	// Por padrão, o resultado será justificado à direita e
	// preenchido com espaços.
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// Também é possível especificar a largura de floats
	// impressos, embora geralmente também seja interessante
	// restringir a precisão decimal ao mesmo tempo com a
	// sintaxe `width.precision`.
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// Para justificar à esquerda, utiliza-se a flag `-`.
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Também pode ser necessário controlar a largura
	// ao formatar strings, especialmente para assegurar
	// que todos os resultados fiquem alinhados, como uma
	// tabela. Para justificar à direita.
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// Para justificar à esquerda, utiliza-se a flag `-` com números.
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Até agora vimos `Printf`, que imprime strings
	// formatadas no `os.Stdout`. `Sprintf` apenas formata
	// e retorna a string sem imprimir em nenhum lugar.
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// É possível formatar e printar com `io.Writers`
	// além de com `os.Stdout` usando `Fprintf`.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
