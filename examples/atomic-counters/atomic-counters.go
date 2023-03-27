// O mecanismo primário para gerenciar estado em Go ẽ
// a comunicação entre canais. Um dos exemplos apresentados
// foi os [worker pools](worker-pools). No entanto, existem
// algumas outras opções para gerenciar estado.
// Aqui será apresentado o pacote `sync/atomic` para _contadores
// atômicos_ accessados por múltiplas goroutines.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Será utilizado um inteiro _unsigned integer_
	// para representar um contador sempre positivo.
	var ops uint64

	// Um WaitGroup auxiliará a aguardar todas as goroutines
	// finalizarem suas tarefas.
	var wg sync.WaitGroup

	// Iniciaremos 50 goroutines que incrementarão o contador
	// exatamente 1000 vezes.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// Para incrementar o contador será utilizado
				// `AddUint64`, passando como parâmetro o
				// endereço de memória do contador `ops`.
				// Note que para acessar o endereço de memória
				// de determinado dado, se utiliza a sintaxe `&`.
				atomic.AddUint64(&ops, 1)
				// ops++
			}
			wg.Done()
		}()
	}

	// Aguarda todas as goroutines finalizarem.
	wg.Wait()

	// É seguro ler o contador `ops` uma vez que não
	// há nenhuma goroutine escrevendo sobre ele.
	// Para ler o contador enquanto ele ẽ atualizado,
	// pode-se utilizar a função `atomic.LoadUint64`,
	// por exemplo.
	fmt.Println("ops:", ops)
}
