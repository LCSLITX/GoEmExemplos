// A biblioteca padrão de Go, oferece um excelente suporte
// para clientes HTTP e servidores no pacote `net/http`.
// Neste exemplo, será usado o pacote para enviar
// simples requisições HTTP.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Envia uma requisição HTTP com método GET para um
	// servidor. `http.Get` é um atalho conveniente para
	// criar um objeto `http.Client` e chamar seu método
	// `Get`; É utilizado o objeto `http.DefaultClient`
	// que tem algumas configurações padrão úteis.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Imprime o status da resposta HTTP.
	fmt.Println("Response status:", resp.Status)

	// Imprime as primeiras 5 linhas do corpo da resposta.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
