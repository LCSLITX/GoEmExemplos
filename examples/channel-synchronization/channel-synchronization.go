// É possível utilizar canais para sincronizar execuções
// de goroutines. Aqui está um exemplo de utilização
// de um recebimento bloqueante para aguardar a goroutine
// finalizar. Ao aguardar por múltiplas goroutines terminarem,
// talvez seja melhor utilizar [WaitGroup](waitgroups).

package main

import (
	"fmt"
	"time"
)

// Esta função será executada em goroutine.
// O canal `done` será usado para notificar
// outra goroutine que esta função foi finalizada.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Envio de valor para notificar que a execução
	// finalizou
	done <- true
}

func main() {

	// Inicializando a _ worker goroutine_, e passando o canal
	// que será notificado.
	done := make(chan bool, 1)
	go worker(done)

	// O código é bloqueado até a notificação ser recebida.
	<-done
}
