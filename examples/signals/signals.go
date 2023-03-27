// Por vezes, é necessário que programas em Go tratem
// [sinais Unix](https://en.wikipedia.org/wiki/Unix_signal).
// Por exemplo, pode-se querer que um servidor seja desligado
// graciosamente ao receber um sinal `SIGTERM`, ou que uma
// ferramenta de linha de comando pare de processar um input
// ao receber um sinal `SIGINT`.
// Aqui está como tratar estes sinais com canais em Go.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// A notificação de sinais em Go, funciona enviando valores
	// do tipo `os.Signal` em um canal. Aqui é criado um canal
	// para receber estas notificações. Note que este canal deve
	// ser `buffered`.
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registra determinado canal para
	// receber notoficações de sinais específicos.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// O sinal poderia ser recebido pelo `sigs` aqui
	// na função main, mas será feito em uma goroutine
	// separada, para demonstrar como pode ser feito o
	// desligamento gracioso de maneira mais realista.
	done := make(chan bool, 1)

	go func() {
		// Esta goroutine executa um recebimento bloqueante
		// de sinais. Ao receber um, o sinal será impresso
		// e então o programa será notificado que pode ser
		// finalizado.
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// O código irá aguardar aqui até que seja recebido
	// o sinal esperado (como indicado pela goroutine
	// acima, que enviará um valor no canal `done`) e
	// então finalizar.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
