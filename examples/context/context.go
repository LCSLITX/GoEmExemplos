// No exemplo anterior foi apresentado como configurar um simples
// [servidor HTTP](http-server). Os servidores HTTP são úteis para
// demonstrar o uso de `context.Context` e controle de cancelamento.
// Um `Contexto` carrega deadlines, sinais de cancelamento
// e outros valores do escopo da requisição através da API
// e goroutines.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// Um `context.Context` é criado para cada request
	// pelo próprio pacote `net/http`, e fica disponível
	// com o método `Context()`.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Aguarda alguns segundos antes de enviar a resposta
	// para o cliente, para simular algum trabalho que o
	// servidor possa fazer. Enquanto trabalha, acompanha
	// o canal `Done()` do Contexto, para esperar por um
	// sinal de cancelamento e retornar a resposta o mais
	// rápido possível.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// O método `Err()` do contexto retorna um erro
		// que explica porque o canal `Done()` foi fechado.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// Como antes, o handler é registrado na rota "/hello"
	// e o servidor é iniciado em seguida.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
