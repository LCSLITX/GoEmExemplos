// Usa-se o `os.Exit` para, imediatamente, sair do programa
// com um determinado status.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `defer` NÃO serão executados quando utilizado `os.Exit`,
	// então este `fmt.Println` nunca será chamado.
	defer fmt.Println("!")

	// Exit com status 3.
	os.Exit(3)
}

// Note que diferentemente de C, por exemplo, Go não
// utiliza um inteiro como valor de retorno de main,
// para indicar status de saída. Se a intenção for sair
// de um programa com status não-zero (sem sucesso),
// `os.Exit` deve ser usado.
