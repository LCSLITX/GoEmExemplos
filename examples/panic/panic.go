// Um `panic` tipicamente significa que acontenceu algum
// erro inesperado. Geralmente é utilizado para falhar
// rapidamente em erros que não deviam ocorrer durante
// uma operação normal, ou que não está sendo tratado
// de maneira _graciosa_ ou apropriada.

package main

import "os"

func main() {

	// O panic será utilizado por todo o site para checar por
	// erros inesperados. Este é o único código em todo o site
	// especificamente desenhado para causar um panic.
	panic("uhhh... Houston, we have a problem.")

	// Um uso comum do panic é para abortar a execução
	// de determinado código se uma função retorna um
	// erro que não é tratado. Aqui está um exemplo
	// simulado de `panic` ao receber um erro inesperado
	// ao se criar um novo arquivo.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
