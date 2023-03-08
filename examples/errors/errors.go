// Em Go, é considerado idiomático comunicar erros via
// retornos separados e explícitos. Isto contrasta com
// exceções usadas em outras linguagens como Java e Ruby
// com um único resultado sobrecarregado; ou com valores
// de erros por vezes usado em C. O _approach_ de Go torna
// muito mais fácil a verificação de que funções estão
// retornando erros e tratá-los usando os mesmos mecanismos
// empregados pela linguagem em outras tarefas não relacionadas
// a erros.

package main

import (
	"errors"
	"fmt"
)

// Por convenção, erros são os últimos valores retornados
// e implementam a interface nativa `erro`.
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` constrói um valor `error` básico
		// com a mensagem de erro determinada.
		return -1, errors.New("can't work with 42")

	}

	// Um valor `nil` como último retorno de uma determinada função
	// indica que não há nenhum erro.
	return arg + 3, nil
}

// É possível usar tipos de erro customizados implementado
// o método `Error()`. Aqui está uma variante do exemplo
// acima que usa um tipo erro customizado para representar
// explicitamente um erro de argumento.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// Neste caso, pode-se utilizar a sintaxe `&argError`
		// para construir uma nova struct, passando valores para
		// os campos `arg` e `prob`.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// Os dois loops abaixo, testam cada uma das funções
	// que retornam erros. Note que a utilização de uma
	// verificação de erro em linha, junto ao `if` é algo
	// comum em Go.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// É possível utilizar programaticamente os dados
	// em um erro customizado. Será preciso capturar o
	// erro como uma instância do tipo de erro
	// customizado via asserção de tipo.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
