// As vezes, é necessário executar código em algum momento
// futuro, ou repetidamente com algum intervalo. Go conta
// com as ferramentas nativas _timer_ e _ticker_. Ambos os
// recursos tornam a tarefa fácil.
// Primeiro será apresentado o recurso timer e posteriormente
// o [ticker](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Timers representam um único evento no futuro.
	// Ao timer é informado quanto tempo se pretende esperar
	// e o recurso providencia um canal que será notificado
	// ao esgotar o tempo estipulado.
	// Este timer irá aguardar 2 segundos.
	timer1 := time.NewTimer(2 * time.Second)

	// O `<-timer1.C` bloqueia no canal `C` de Timer
	// até que envie um valor indicando que o tempo
	// foi atingido.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Se a intenção for apenas aguardar, é possível utilizar
	// `time.Sleep`. Uma razão para o timer ser útil é que ele
	// pode ser cancelado antes o tempo esgotar.
	// Aqui está um exemplo.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Ao dar tempo suficiente para o `timer2` ser disparado,
	// é fácil concluir que ele foi, de fato, parado, pois
	// não disparou.
	time.Sleep(2 * time.Second)
}
