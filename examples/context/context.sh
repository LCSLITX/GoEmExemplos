# Executa o servidor em background.
$ go run context-in-http-servers.go &

# Simula uma requisição de cliente para `/hello`,
# e logo após, utiliza Ctrl+C para enviar sinal
# de cancelamento.
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
