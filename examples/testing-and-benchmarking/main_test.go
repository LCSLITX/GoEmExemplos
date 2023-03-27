// Testes unitários é uma parte importante de se
// escrever programas em Go. O pacote `testing`
// fornece as ferramentas necessárias para escrever
// testes unitários e o comando `go test` executa
// os tests.

// Pela didática da demonstração, este código está no pacote
// `main`, mas poderia estar em qualquer pacote. Tipicamente,
// os testes permanecem no mesmo pacote em que o código
// principal está.
package main

import (
	"fmt"
	"testing"
)

// Será testada uma implementação simples de uma função
// que verifica qual inteiro é menor. Usualmente,
// o código a ser testado ficaria isolado num arquivo
// com nome similar a `intutils.go`, e outro arquivo apenas
// para testes, teria o nome parecido com `intutils_test.go`.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Um teste é criado ao escrever uma função com a
// palavra `Test` no início do nome.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` will report test failures but continue
		// executing the test. `t.Fatal*` will report test
		// failures and stop the test immediately.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Escrever testes pode ser algo repetitivo, então é idiomático
// usar *table-driven style* ou *estilo orientado a tabela*,
// em que os inputs de teste e os outputs esperados são listados
// numa tabela e um simples loop, interage com a tabela e realiza
// os testes.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// t.Run permite a execução de "subtests",
		// um para cada entrada da tabela. Estes são
		// exibidos separadamente ao executar o comando
		// `go test -v`.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Testes de benchmark também ficam em arquivos `_test.go`
// e são nomeados com a palavra `Benchmark` no início.
// A execução do `testing` executa cada função de benchmark várias
// vezes, aumentando `b.N` em cada execução até que seja realizada
// uma medida precisa.
func BenchmarkIntMin(b *testing.B) {
	// Tipicamente o benchmark roda a função que
	// está a ser avaliada em um loop, por `b.N` vezes.
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
