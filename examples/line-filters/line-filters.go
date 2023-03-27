// Um _line filter_, ou filtro de linha, é um tipo comum
// de software que lê o input em `stdin`, processa,
// e então imprime um resultado em `stdout`.
// Os comandos `grep` e `sed` do Linux são exemplos
// de line filters.

// Aqui está um exemplo de filtro de linha em Go,
// que escreve uma versão em letras maiusculas do
// texto de entrada. Este padrão pode ser utilizado
// para escrever novos filtros de linha.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Envelopar o `os.Stdin` `unbuffered` com um
	// scanner `buffered` resulta em um método
	// conveniente de `Scan` que avança o scanner para
	// o próximo token; que é a próxima linha em um scanner
	// padrão.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` retorna o token atual, a próxima
		// linha do input.
		ucl := strings.ToUpper(scanner.Text())

		// Escreve a linha com letras maiusculas.
		fmt.Println(ucl)
	}

	// Verfica erros durante o `Scan`. O final do arquivo
	// é esperado pelo `Scan` e não reportado como um erro.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
