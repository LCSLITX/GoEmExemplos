// A ferramenta _select_ permite aguardar múltiplos
// canais. Combinar goroutines e canais com select,
// é um recurso muito útil em Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Para este exemplo, o select será utilizado com
	// dois canais.
	c1 := make(chan string)
	c2 := make(chan string)

	// Cada canal receberá um valor depois de certo tempo
	// para simular, por exemplo, operações RPC bloqueantes
	// executando em goroutines concorrentes.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "um"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "dois"
	}()

	// O `select` será utilizado para aguardar ambos os valores
	// simultaneamente, imprimindo cada um conforme são recebidos.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recebido", msg1)
		case msg2 := <-c2:
			fmt.Println("recebido", msg2)
		}
	}
}
