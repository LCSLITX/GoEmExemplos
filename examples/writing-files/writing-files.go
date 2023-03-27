// Escrever arquivos em Go segue o mesmo padrão
// apresentado para os casos de leitura.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Para começar, aqui está como escrever uma string
	// (ou apenas bytes) em um arquivo.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Para escritas mais granulares, se abre o arquivo.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// É idiomático utilizar defer para fechar um arquivo
	// com `Close` imediatamente após abri-lo.
	defer f.Close()

	// É possível escrever `Write` slices de bytes.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// O método `WriteString` também está disponível.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Usa-se o `Sync` para descarregar a escrita
	// para um armazenamento estável.
	f.Sync()

	// O pacote `bufio` também fornece escritores,
	// ou writers, `buffered` em adição aos leitores
	// anteriormente vistos.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Usa-se o `Flush` para assegurar que todas as
	// operações `buffered` foram aplicadas ao escritor.
	w.Flush()
}
