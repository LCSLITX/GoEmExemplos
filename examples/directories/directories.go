// Go possui várias funções úteis para trabalhar
// com diretórios em um sistema de arquivos.

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

	// Cria um novo sub-diretório no atual diretório.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Ao criar diretórios temporários, é uma boa prática
	// adicionar um `defer` para removê-los. `os.RemoveAll`
	// irá deletar uma árvore de diretórios por completo
	// (similar ao comando linux rm -rf).
	defer os.RemoveAll("subdir")

	// Função auxiliar para criar um novo arquivo vazio.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// É possível criar hierarquia de diretórios, inclusive
	// parentes com `MkdirAll`. Isto é similar ao comando de
	// linux `mkdir -p`.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` lista o conteúdo do diretório,
	// retornando uma slice de objetos do tipo
	// `os.DirEntry`.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` permite mudar o diretório de trabalho atual,
	// similar ao comando de linux `cd`.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Agora, será visto a lista de conteúdos de
	// `subdir/parent/child` quando listar o
	// diretório *atual* (".").
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `cd` de volta ao ponto de início.
	err = os.Chdir("../../..")
	check(err)

	// É possível visitar um diretório *recursivamente*,
	// incluindo todos os seus sub-diretórios. `Walk`
	// aceita uma função callback para tratar todo arquivo
	// ou diretório visitado.
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// `visit` é chamado recursivamente para cada arquivo
// e diretório encontrado por `filepath.Walk`.
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
