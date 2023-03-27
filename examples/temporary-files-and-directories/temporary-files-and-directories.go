// Muitas vezes, ao longo da execução de um programa, é preciso
// criar dados que não são necessários após seu encerramento.
// *Arquivos e diretórios temporários* são úteis para esta
// finalidade, pois não poluem o sistema de arquivos
// permanentemente.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// A maneira mais fácil de criar um arquivo temporário é
	// com `os.CreateTemp`. eta função cria um arquivo e
	// abre para leitura e escrita. Ao fornecer `""` como o
	// primeiro argumento, a função irá criar o arquivo no
	// local padrão do sistema operacional.
	f, err := os.CreateTemp("", "sample")
	check(err)

	// Este trecho exibe o nome do arquivo temporário.
	// Em Sistemas Operacionais baseados em Unix, o
	// diretório em que o artquivo será criado, provavelmente,
	// é o `/tmp`. O nome do arquivo começa com o prefixo
	// passado como segundo argumento para `os.CreateTemp` e o
	// resto é escolhido automaticamente para garantir que
	// eventuais chamadas concorrentes sempre criarão nomes
	// de arquivos diferentes.
	fmt.Println("Temp file name:", f.Name())

	// Limpa o arquivo depois de finalizada a execução.
	// É provável que o Sistema Operacional tenha uma
	// rotina para limpar arquivos temporários
	// esporadicamente, mas é uma boa prática fazer isso
	// explicitamente ao final do código.
	defer os.Remove(f.Name())

	// É possível escrever dados no arquivo.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Se a intenção for escrever muitos arquivos
	// temporários, é preferível criar um diretório
	// temporário. A função `os.MkdirTemp` recebe os
	// mesmos argumentos que a `CreateTemp`, mas
	// retorna um nome de diretório, ao invés de
	// um arquivo aberto.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Agora é possível sintetizar os nomes de arquivos
	// temporários ao incluir como prefixo o diretório
	// temporário.
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
