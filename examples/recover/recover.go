// Go possibilita recuperar, ou _recover_ de um panic,
// utilizando a função nativa `recover`. O `recover`
// pode interromper um `panic` de abortar a execução
// de um determinado código e prosseguir normalmente.

// Um exemplo em que isto pode ser útil: um servidor
// não deveria quebrar se uma conexão com um cliente causa
// um erro crítico. Ao invés disso, o mais apropriado seria
// fechar a conexão e continuar servindo outros clientes.
// Na verdade, isso é o que o pacote `net/http` nativo de Go,
// faz por padrão para servidores HTTP.

package main

import "fmt"

// Esta função gera um panic.
func mayPanic() {
	panic("uhhh... Houston?")
}

func main() {
	// `recover` deve ser chamado dentro de uma função deferida.
	// Quando uma função gerar panico, o defer será ativado e a
	// função recover será chamada para capturar o panic.
	defer func() {
		if r := recover(); r != nil {
			// O valor retornado pela função `recover`
			// é o erro gerado na chamada do `panic`.
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// Este código não será executado, porque `mayPanic`
	// gera panico. A execução de main é interrompida
	// no momento do pânico e é resumida no fechamento
	// deferido.
	fmt.Println("After mayPanic()")
}
