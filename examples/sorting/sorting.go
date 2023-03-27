// O pacote `sort` de Go implementa ordenação tanto para
// tipos nativos como para tipos definidos pelo usuário.
// Aqui, será apresentado sorteamento para tipos nativos.

package main

import (
	"fmt"
	"sort"
)

func main() {

	// Métodos de `sort` são específicos de tipos nativos;
	// aqui está um exemplo para slices de strings. Note que
	// a ordenação modifica o original, então atualiza o valor
	// da slice e NÃO retorna nova estrutura.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	// Um exemplo de ordenação em slices de `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:    ", ints)

	// Também é possível utilizar o `sort` para checar se
	// uma slice já está ordenada.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Ordenado:", s)
}
