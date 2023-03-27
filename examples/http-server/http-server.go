// Criar um servidor HTTP básico é bem simples
// utilizando o pacote `net/http`.

package main

import (
	"fmt"
	"net/http"
)

// Um conceito fundamental nos servidores do pacote `net/http`
// é o de *handlers*. Um handler é um objeto implementando
// a interface `http.Handler`. Uma forma comum de escrever
// um handler é usando o adaptador `http.HandlerFunc`
// em funções com a assinatura apropriada.
func hello(w http.ResponseWriter, req *http.Request) {

	// Funções que servem como handlers recebem dois
	// parâmetrods, primeiro o `http.ResponseWriter` e
	// segundo o `http.Request`. O _response writer_ é
	// usado para escrever a resposta HTTP.
	// Aqui um exemplo de resposta com um simples
	// "hello\n".
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// Este handler faz algo um pouco mais sofisticado,
	//lendo todo o header, ou cabeçalho, da requisição
	// HTTP e repetindo no corpo da resposta.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Os handlers são registrados em rotas do servidor
	// usando função `http.HandleFunc` que configura o
	// *roteador padrão* no pacote `net/http` e recebe uma função
	// como argumento.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Finalmente, é chamada a função `ListenAndServe`
	// com indicação da porta em que o servidor será
	// servido e o handler. Passando `nil` como handler,
	// faz com que o roteador padrão seja utilizado.
	http.ListenAndServe(":8090", nil)
}
