// _Timeouts_ são importantes para programas que conectam
// em recursos externos ou que de outra forma precisam
// limitar o tempo de execução. A implementação em Go
// é fácil e elegante graças aos recursos `channels` e
// `select`.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Para este exemplo, suponha que se está executando uma
	// chamada externa que retorna um resultado no canal `c1`
	// após 2 segundos. Note que o canal é `buffered`, então o
	// `send` na goroutine é _não bloqueante_ ou _nonblocking_,
	// este é um padrão comum para prevenir vazamento de goroutines
	// no caso do canal nunca ser lido.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Aqui está o `select` implementando o timeout.
	// `res := <-c1` aguarda o resultado e `<-time.After`
	// aguarda o valor a ser enviado depois do timeout de
	// 1 segundo. Como o `select` prossegue com o primeiro
	// recebimento que está pronto, será executado o caso com
	// timeout se a operação levar mais do que o tempo permitido
	// de 1 segundo.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Se for permitido um timeout mais de longo de 3 segundos,
	// então o recebimento de c2 será executado e o valor impresso.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
