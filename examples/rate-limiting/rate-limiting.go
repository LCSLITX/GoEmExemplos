// [_Rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting)
// ou limitador de frequência é um mecanismo importante para
// controlar a utilização de recursos e manter a qualidade
// do serviço. Go suporta o rate limiting com goroutines,
// canais e [tickers](tickers) de maneira bem elegante.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Primeiro será apresentado rate limiting básico.
	// Suponha que seja necessário limitar o tratamento
	// das requisições recebidas, que serão servidas em
	// um canal `requests`.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Este canal limitador, ou `limiter`, receberá um valor
	// a cada 200 milisegundos. Este é o regulador no esquema
	// de rate limiting.
	limiter := time.Tick(200 * time.Millisecond)

	// Ao bloquear um recebimento do canal limitador
	// antes de atender a cada solicitação,
	// será tratada 1 solicitação por vez,
	// a cada 200 milissegundos.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Também é possível permitir envio de lotes de pequena
	// quantidade ou `short bursts` (rajadas curtas) de requisições
	// no esquema de rate limiting enquanto se cumpre com a
	// limitação de frequência no geral.
	// Para fazer isto, é necessário utilizar buffer no canal
	// `limiter`. Este canal `burstyLimiter` permitirá envios
	// de lotes de até 3 requisições.
	burstyLimiter := make(chan time.Time, 3)

	// Preenchendo o canal com requisições para representar
	// os lotes, ou bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// A cada 200 milisegundos, será adicionado um valor
	// ao canal `burstyLimiter`, até o limite de 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Agora é simulado outras 5 requisições sendo recebidas.
	// As 3 primeiras serão beneficiadas pela capacidade do
	//`burstyLimiter`de aceitar requisições em lote .
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
