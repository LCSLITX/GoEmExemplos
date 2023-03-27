// Para aguardar múltiplas goroutines finalizarem,
// é possível utilizar waitgroups.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Essa função será executada em cada goroutines.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Será utilizado `sleep` para simular uma
	// tarefa extensa.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// Este WaitGroup é usado para esperar por todas as
	// goroutines lançadas finalizarem.
	// Note: se um WaitGroup é explicitamente passado
	// para uma função como parâmetro, deve ser feito com
	// *ponteiro*.
	var wg sync.WaitGroup

	// Inicia vãrias goroutines e incrementa o contador do
	// WaitGroup para cada uma.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// Evite a utilização do mesmo valor de índice em
		// cada goroutine closure.
		// veja a [FAQ](https://golang.org/doc/faq#closures_and_goroutines)
		// para mais detalhes.
		i := i

		// Aqui se envelopa a chamada do worker em uma closure
		// para informar ao WaitGroup que o worker finalizou a
		// tarefa. Desta forma o próprio worker não precisa saber
		// da concorrência envolvendo sua execução.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Bloqueado até que o contador do WaitGroup volte a 0;
	// e que todos os workers sejam notificados que eles
	// finalizaram.
	wg.Wait()

	// Note que o código desta forma não possibilita nenhuma
	// forma direta para propagar erros dos workers.
	// Para mais casos avançados de uso, considere usar o pacote
	// [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup).
}
