// Neste exemplo, será apresentado como implementar
// uma _worker pool_ usando goroutines e canais.

package main

import (
	"fmt"
	"time"
)

// Aqui está um worker, a partir do qual executará várias
// instâncias concorrentes. Estes workers receberão tarefas
// pelo canal `jobs` a enviarão o resultado correspondente
// pelo canal `results`. Será utilizado `time.Sleep` para
// simular uma tarefa extensa.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// Para usar os workers, é preciso enviar as tarefas
	// a serem executadas e coletar os resultados, para
	// isso são criados os dois canais com `make`.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Isto inicia 3 workers, inicialmente bloqueados porque
	// ainda não há tarefas.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Aqui são enviadas 5 tarefas e então o canal é fechado
	// para indicar que essas são todas as tarefas a serem
	// executadas.

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finalmente são coletados todos os resultados das tarefas.
	// Isso também assegura que todas as goroutines dos workers
	// estão finalizadas. Um método alternativo de aguardar
	// várias goroutines é com [WaitGroup](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
