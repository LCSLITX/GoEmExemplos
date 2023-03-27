// O pacote `filepath` fornece funções para parsear
// e construir *file paths*, ou caminhos de arquivos,
// de uma maneira flexível que pode ser utilizado
// em vários sistemas operacionais; `dir/file` no
// Linux vs. `dir\file` no Windows, por exemplo.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` deve ser usado para construir caminhos
	// de forma flexível. Pode receber qualquer número
	// de argumentos e constrói um caminho com hierarquia
	// a partir deles.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Sempre deve ser utilizado o `Join` ao invés de
	// concatenar `/` ou `\` manualmente. Em adição
	// a a tornar flexível e utilizável em mais de um
	// sistema operacional, `Join` também irá normalizar
	// os caminhos, removendo separadores superfluos e
	// mudanças de diretório.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` e `Base` podem se rusados para dividir o caminho
	// para o diretório e para o arquivo. Alternativamente,
	// `Split` retornará ambos em uma mesma chamada.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// É possível checar se um caminho é absoluto.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Alguns nomes de arquivos tem extenção seguida
	// de um ponto. É possível separar a extensão do
	// resto do nome com `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Para encontrar nomes de arquivos com extensão
	// usa-se `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` encontra o caminho relativo entre *base* e
	// *alvo*, ou *target*. Retorna erro se o alvo não
	// pode ser acessar a partir de base.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
