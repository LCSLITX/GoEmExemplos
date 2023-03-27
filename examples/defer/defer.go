// _Defer_ é utilizado para assegurar que uma chamada de função
// é realizada tardiamente numa execução de código, geralmente
// para limpeza de memória. `defer` é frequentemente utilizada
// em situações que, por exemplo, `ensure` e `finally`
// seriam usadas em outras linguagens.

package main

import (
	"fmt"
	"os"
)

// Suponha que quiséssemos criar um arquivo, escrever
// nele e então fechá-lo quando tivéssemos terminado.
// Aqui está um exemplo de como isso pode ser feito
// utilizando `defer`.
func main() {

	// Imediatamente depois de abrir um arquivo com
	// `createFile`, é colocado o `defer` para, literalmente,
	// _adiar_ o fechamento daquele arquivo com `closeFile`.
	// Isto será executado ao final da função que envolve a
	// chamada, no caso a `main`, depois que `writeFile`
	// for finalizado.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// É importante verificar erros ao fechar um arquivo,
	// mesmo numa função adiada ou `defer`ida.
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
