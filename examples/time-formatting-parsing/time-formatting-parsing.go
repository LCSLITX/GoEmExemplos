// Go suporta formatação de tempo e parseamento
// via layouts baseados em padrões.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Aqui está um exemplo básico de formatação de tempo
	// de acordo com RFC3339, usando a constante com layout
	// correspondente.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// O parseamento de tempo usa o mesmo valor de layout que `Format`.
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` e `Parse` usam layouts baseados em exemplo. Geralmente
	// serão utilizadas constantes de `time` para estes layouts, mas
	// também é possível fornecer layouts personalizados. Os Layouts precisam
	// usar o tempo referência `Mon Jan 2 15:04:05 MST 2006` para exibir o
	// padrão com o qual será formatado/parseado um determinado tempo/string.
	// O exemplo de tempo deve ser exatamente como mostrado: o ano 2006,
	// 15 horas, Monday (segunda-feira) para o dia da semana, etc.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// Para representações puramente numericas, também é
	// possível usar o padrão string, com os componentes
	// extraídos do valor `tempo`.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` retornará um erro caso a entrada esteja
	// mal formada, explicando o problema ao parsear.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
