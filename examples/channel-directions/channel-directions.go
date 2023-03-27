// Ao utilizar canais como parâmetros de funções, é possível
// especificar se o canal servirá apenas para enviar ou receber
// valores. Esta especificidade é uma boa prática pois aumenta a
// tipagem, ou _type-safety_ do código.

package main

import "fmt"

// Esta função `ping` apenas aceita um canal para
// enviar valores. Ao tentar receber valores neste
// canal, resultaria em erro de compilação.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Esta função `pong` aceita um canal para receber
// (`pings`) e um segundo canal para enviar (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
