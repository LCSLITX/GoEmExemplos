// Go oferece suporte nativo para criação dinâmica de
// conteúdo ou exibição de saída personalizada ao usuário
// com o pacote `text/template`. Um outro pacote
// `html/template` fornece a mesma API but tem recursos de
// segurança adicional e deve ser usado para gerar HTML.

package main

import (
	"os"
	"text/template"
)

func main() {

	// É possível criar um novo template e parsear o body
	// a partir de uma string. Templates são uma mistura de
	// texto estático e "ações" envelopadas em `{{...}}` que
	// são usadas dinamicamente para inserir conteúdos.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Alternativamente, é possível utilizar a função
	// `template.Must` para gerar panic no caso do `Parse`
	// returnar um erro. Isto é útil especialmente para
	// templates inicializados em escopo global.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// Ao "executar" um template, é gerado seu texto com alguns
	// valores específicos para suas "ações". A ação `{{.}}` é
	// substituída pelo valor passado como parâmetro para a
	// função `Execute`.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Função auxiliar que será usada abaixo.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Se o dado é uma struct é possível usar a ação
	//`{{.FieldName}}` para acessar os campos. Os campos
	// devem ser exportados para serem acessíveis quando
	// um template está sendo executado.
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// O mesmo se aplica para maps; mas aqui não há restrição
	// de nomes das chaves.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// É possível utilizar if/else para condicionar a
	// execução de templates. Um valor é considerado
	// false se é o valor padrão de um tipo, como 0,
	// ou uma string vazia, nil pointer, etc.
	// Esta amostra demonstra outro recurso dos templates:
	// usando `-` em ações para eliminar espaços vazios.
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// Blocos de `range` permitem iterar sobre slices,
	// arrays, maps ou canais. Dentro de um bloco de
	// range `{{.}}` representa o item atual da iteração.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
