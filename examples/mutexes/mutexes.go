// No exemplo anterior, foi apresentado como gerenciar
// estados de contadores simples usando
// [atomic operations](atomic-counters).
// Para estados mais complexos é recomendado utilizar
// [_mutex_](https://en.wikipedia.org/wiki/Mutual_exclusion)
// para acessar dados entre múltiplas goroutines seguramente.
// MutEx é abreviação de _Mutual Exclusion_

package main

import (
	"fmt"
	"sync"
)

// A struct container possui um map de contadores;
// como a intenção aqui é atualizar concorrentemente
// a partir de múltiplas goroutines, é adicionado `Mutex`
// para sincronizar o acesso.
// Note que mutex não deve ser copiado, então, se a struct
// é passada para outras funções ou etc, deve ser feito
// com ponteiro.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Trava o mutex antes de acessar `counters`; destrava
	// ao final da função utilizando a expressão [defer](defer).
	// e incrementa determinado contador.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// Note que o valor padrão, ou zero, de um mutex é utilizável,
		// então não é necessária a inicialização aqui.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Esta função incrementa, em um loop, um contador de
	// determinado nome.
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// Executando várias goroutines concorrentemente;
	// Note que todas tem acesso ao mesmo `Container`,
	// e duas delas tem acesso ao mesmo contador `a`.
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// Aguarda todas as goroutines finalizarem.
	wg.Wait()
	fmt.Println(c.counters)
}
