// Fechar um canal indica que mais nenhum valor
// será enviado para ele. Isto pode ser útil para
// comunicar a finalização dos recebimentos.

package main

import "fmt"

// Neste exemplo será usado um canal `tarefas` para
// comunicar tarefa a ser executada pela função `main()`
// para uma worker goroutine. Ao não ter mais tarefas
// para a worker, o canal será fechado com `close`.
func main() {
	tarefas := make(chan int, 5)
	pronto := make(chan bool)

	// Aqui está a worker goroutine. Ela repetidamente recebe
	// do canal `tarefas` com `t, mais := <-tarefas`. Nesta forma de
	// recebimento de dois valores, o valor `mais` será falso se
	// `tarefas` foi fechado e todos os valores do canal já
	// tiverem sido recebidos.
	// Isto é utilizado para notificar ao `done` quando todas as
	// tarefas já foram executadas.
	// Importante notar que `tarefas` é um canal `buffered`.
	go func() {
		for {
			t, mais := <-tarefas
			if mais {
				fmt.Println("received job", t)
			} else {
				fmt.Println("received all tarefas")
				pronto <- true
				return
			}
		}
	}()

	// Isto envia 3 tarefas para a worker goroutine pelo canal
	// `tarefas` e então o fecha.
	for t := 1; t <= 3; t++ {
		tarefas <- t
		fmt.Println("sent job", t)
	}
	close(tarefas)
	fmt.Println("sent all tarefas")

	// A execução é aguardada utilizando a
	// [sincronização](channel-synchronization),
	// como apresentado anteriormente.
	<-pronto
}
