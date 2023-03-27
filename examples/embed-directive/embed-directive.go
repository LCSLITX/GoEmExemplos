// `//go:embed` é uma diretiva de compilador ou [compiler
// directive](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives) que
// permite programas a inclusão de arquivos e diretórios arbitrários ao
// binário de Go no momento de build. Leia mais sobre embed directive
// [aqui](https://pkg.go.dev/embed).
package main

// Importação do pacote `embed`; Se não for utilizado
// nenhum identificador exportado por esse pacote, é
// possível usar um import em branco com`_ "embed"`.
import (
	"embed"
)

// `embed` directives aceitam caminhos relativos para o
// diretório contendo o código fonte de um programa em
// Go. Esta diretiva _embeds_ ou _incorpora_  o conteúdo
// do arquivo numa variável `string` imediatamente após
// a diretiva.
//
//go:embed folder/single_file.txt
var fileString string

// Ou incorporam os conteúdos do arquivo em um `[]byte`.
//
//go:embed folder/single_file.txt
var fileByte []byte

// Também é possível incorporar múltiplos arquivos ou
// até mesmo diretórios com `wildcards`. Este utiliza
// a variável do tipo [embed.FS](https://pkg.go.dev/embed#FS),
// o qual implementa um simples sistema de arquivo virtual.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Imprime o conteúdo de `single_file.txt`.
	print(fileString)
	print(string(fileByte))

	// Recupera alguns arquivos do diretório incorporado.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
