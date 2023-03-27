// Go suporta
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>funções recursivas</em></a>.

package main

import "fmt"

// Esta função `fact` chama a si mesma várias vezes,
// decrementando o parâmetro passado a cada chamada,
// até que atinja o caso base `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Closures também podem ser recursivos, mas isso requer que
	// sejam declarados explicitamente com uma variável tipada `var`
	// antes de ser definido.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// `fib` foi anteriormente declarada dentro de `main`,
		// então será executada normalmente.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
