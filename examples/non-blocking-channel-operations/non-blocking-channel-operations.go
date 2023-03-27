// Envios e recebimentos básicos nos canais, são bloqueantes.
// No entanto, é possível usar o `select` com uma cláusula `default`
// para implementar envio e recebimentos _non-blocking_ e até
// non-blocking `select` de multiplos casos.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Aqui está um recebimento não bloqueante. Se um valor
	// estiver disponível em `messages` o `select` prosseguirá
	// com o `case` que tem o `<-messages`. Se não,
	// presseguirá com o case padrão ou `default`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Um send não bloqueante funciona de maneira similar.
	// Aqui, `msg` não pode ser enviado para o canal `messages`,
	// justamente porque o canal não tem buffer e
	// não tem um recebimento. Portanto, o case `default`
	// será executado.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// É possível utilizar multiplos `case`s acima da clausula
	// `default`. Aqui está uma tentativa de recebimento não bloqueante
	// em ambos os canais  `messages` and `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
