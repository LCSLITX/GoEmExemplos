// Go suporta <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">ponteiros</a></em>,
// permitindo passar referências a valores
// dentro do programa.

package main

import "fmt"

// Será apresentado como ponteiros funcionam em contraste a valores,
// com duas funcões: `zeroval` and `zeroptr`. `zeroval` recebe um
// `int` como parâmetro, então argumentos serão passados para dentro
// da função como um valor. `zeroval` terá uma cópia de `ival`
// distinta da que foi passada no chamamento da função.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` de outra forma, recebe um `*int` como parâmetro,
// ou seja, um ponteiro de `int`. A expressão `*iptr`
// no corpo da função _dereferencia_ o ponteiro de seu endereço
// de memória para o atual valor presente naquele endereço.
// Atribuindo um valor 0 ao ponteiro dereferenciado altera o valor
// presente no endereço.
// Em termos mais claros, diferente do `zeroval`, a função `zeroptr`
// não recebe uma cópia do valor, mas o próprio endereço de memória
// que o valor ocupa. Ao _dereferenciar_ o endereço e atribuir um novo
// valor, se está alterando o próprio valor e não uma cópia.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// A sintaxe `&i` se refere ao `endereço de memória` de `i`,
	// ou seja, ponteiro para `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Ponteiros também podem ser impressos.
	fmt.Println("pointer:", &i)
}
