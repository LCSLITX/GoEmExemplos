// [_Command-line flags_](https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// ou bandeiras de linha de comando, são uma forma
// comum de especificar opções para programas em
// linha de comando. Por exmeplo, em `wc -l` o
//`-l` é uma flag.

package main

// Go fornece um pacote `flag` que suporta parseamento
// básico de flags. Será usado este pacote para a
// implementação do exemplo.
import (
	"flag"
	"fmt"
)

func main() {

	// Declarações básicas de flag estão disponíveis para
	// strings, inteiros e booleanos. Aqui, é declarado uma
	// string com a flag `word` com um valor padrão `"foo"`
	// e uma descrição curta. Esta função `flag.String` retorna
	// um ponteiro de string (não um valor em string);
	// mais adiante será abordado como utilizar este ponteiro.
	wordPtr := flag.String("word", "foo", "a string")

	// Aqui são declaradas as flags `numb` e `fork`, usando
	// um approach similar à flag `word`.
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// Também é possível declarar uma opção que use uma
	// variável existente, declarada em algum lugar do
	// código. Note que será preciso passar um ponteiro
	// para a função que declara a flag.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Uma vez que todas as flags estão declaradas, chama-se
	// `flag.Parse()` para executar o parseamento.
	flag.Parse()

	// Aqui as flags e qualquer eventual argumento posicional
	// serão apenas exibidos. Note que é necessário realizar
	// a _dereference_ dos ponteiros, por exemplo, com a sintaxe
	// `*wordPtr` para acessar os valores.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
