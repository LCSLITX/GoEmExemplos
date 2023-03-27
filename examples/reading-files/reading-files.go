// Ler e escrever arquivos são tarefas básicas que podem ser
// necessárias em muitos programas.
// Primeiro será apresentado como ler arquivos.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Ler arquivos requer checagem de erros.
// Esta função auxiliar ajudará a realizar a verificação.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// A forma mais simples de ler um arquivo é capturar
	// o seu conteúdo e deixa-lo disponível em memória.
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// Geralmente é necessário ter mais controle sobre como
	// e quais partes do arquivo são lidas. Para estas situações,
	// inicia-se abrindo um arquivo para obter o valor `os.File`.
	f, err := os.Open("/tmp/dat")
	check(err)

	// Este trecho lê alguns bytes do início do arquivo.
	// Permite que até 5 sejam lidos, mas é importante
	// notar quantos foram lidos de fato.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Também é posssível procurar, `Seek`, uma localização
	// no arquivo e iniciar a leitura a partir dela.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// O pacote `io` fornece algumas funções que podem
	// ser úteis para a leitura de arquivos. Por exemplo,
	// leituras como as acima, podem ser implementadas
	// de maneira bem robusta com `ReadAtLeast`.
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Não há nenhuma forma nativa de rebobinar, mas
	// `Seek(0, 0)` atinge esse objetivo.
	_, err = f.Seek(0, 0)
	check(err)

	// O pacote `bufio` implementa um leitor `buffered`
	// que pode ser útil tanto para eficiência quanto
	// para muitas leituras pequenas por conta dos
	// métodos de leitura adicionais que são fornecidos.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// É necessário fechar os arquivos com `Close`
	// ao terminar de manipulá-lo (geralmente isso
	// é feito imediatamente depois de abrir, com	o
	// `defer`).
	f.Close()
}
