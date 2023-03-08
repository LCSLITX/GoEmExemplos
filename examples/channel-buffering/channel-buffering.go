// Por padrão canais são _unbuffered_, ou seja, apenas
// aceitam envios (`chan <-`) se há um recebimento
// correspondente (`<- chan`) pronto para receber.
// Canais _Buffered_ aceitam um número limitado de valores
// sem um recebimento correspondente.

package main

import "fmt"

func main() {

	// Aqui é criado um canal de strings que armazena até
	// dois valores, ou seja, trata-se de um canal _buffered_.
	messages := make(chan string, 2)

	// Por ser _buffered_ é possível enviar valores para o
	// canal, sem ter um recebimento correspondente.
	messages <- "buffered"
	messages <- "channel"

	// Posteriormente os valores são recebidos normalmente.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
