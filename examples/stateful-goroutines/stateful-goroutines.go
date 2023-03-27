// No exemplo anterior foi utilizado travamento explícito com
// [mutexes](mutexes) para sincronizar acesso compartilhado a
// estados entre múltiplas goroutines. Outra opção é utilizar
// recursos de sincronização nativa das goroutines e canais
// para atingir o mesmo objetivo. Esta forma baseada em canais
// está alinhada com as ideias de comunicação através do
// compartilhamento de memória de Go, de forma que cada dado
// seja acessado por exatamente uma goroutine.
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Neste exemplo o estado será pertencente a uma única
// goroutine (proprietária). Isso garante que o dado nunca
// seja corrompido com acesso concorrente. Para ler ou escrever
// neste estado, outras goroutines enviarão requisições para a
// goroutine proprietária e receberão as respostas correspondentes.
// As structs `readOp` e `writeOp` encapsulam estas requisições
// e uma forma para a goroutine proprietária responder.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Serão contadas quantas operações são realizadas.
	var readOps uint64
	var writeOps uint64

	// Os canais `reads` e `writes` serão utilizados por outras
	// goroutines para emitir requisições de leitura e
	// escrita respectivamente.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Aqui está a goroutine que possui o estado, que é um
	// map, como no exemplo anterior, mas agora privado à
	// _stateful goroutine_. Esta goroutine possui um
	// `select` que, repetidamente verifica os canais `reads`
	// e `writes`, e responde aos requests conforme são
	// recebidos. A execução acontece primeiro realizando
	// a operação solicitada e então enviando o valor no canal
	// de respostas `resp` para indicar sucesso (e o valor desejado
	// no caso de ser uma operação de leitura `reads`).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Este trecho inicia 100 goroutines que solicitam leituras
	// para a goroutine proprietária do estado, via canal `reads`.
	// Cada leitura requer a construção da struct `readOp`, bem como
	// envio da operação pelo canal `reads` e, então, receber
	// o resultado pelo canal `resp`.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Aqui são iniciadas 10 escritas de forma similar
	// à leitura.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// O time.Sleep serve apenas para deixar
	// as goroutines trabalharem por um segundo.
	time.Sleep(time.Second)

	// Finalmente, as operações realizadas são capturadas
	// e a contagem reportada.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
