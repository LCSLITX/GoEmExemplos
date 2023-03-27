// [Variáveis de ambiente](https://pt.wikipedia.org/wiki/Vari%C3%A1vel_de_ambiente)
// são um mecanismo universal para injetar informações
// e configurações em programas unix.
// Será apresentado como setar, acessar
// e listar variáveis de ambiente.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Para configurar um par de chave/valor,
	// usa-se `os.Setenv`. Para acessar o valor
	// para uma chave, usa-se `os.Getenv`.
	// O segundo, retorna uma string vazia se a chave
	// não for encontrada no ambiente.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// É utilizado `os.Environ` para listar todos os pares
	// de chave/valor presentes no ambiente. Isto retorna
	// uma slice de strings no formato `CHAVE=valor`.
	// É possível utilizar `strings.SplitN` para acessar
	// chaves e valores. Aqui são printadas todas as chaves.
	// Para acesso aos valores seria usado `pair[1]`.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
