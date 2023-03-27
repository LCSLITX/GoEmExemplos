// Em um exemplo [anterior](range) foi apresentado como `for` e
// `range` podem iterar estruturas de dados básicas.
// Também é possível usar esta sintaxe para iterar sobre
// valores recebidos em um canal.

package main

import "fmt"

func main() {

	// Vamos iterar sobre 2 valores no canal `fila`.
	fila := make(chan string, 2)
	fila <- "one"
	fila <- "two"
	close(fila)

	// Este `range` itera sobre cada um dos elementos
	// na medida que são recebidos de `fila`.
	// Como o canal é fechado, com `close`, a iteração finaliza
	// depois de recebidos os dois elementos.
	for elem := range fila {
		fmt.Println(elem)
	}
}
