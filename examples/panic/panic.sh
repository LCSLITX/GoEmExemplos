# Ao executar este código ocorrerá um panic e será
# exibido no console uma mensagem de erro, um 
# "rastro" de execução da goroutine, e saída, 
# ou exit, com código diferente de zero.

# Quando o primeiro panic na função main dispara, 
# o código é interrompido, sem continuar até o 
# final da execução. Para visualizar o que ocorre
# no segundo panic deste exemplo, é necessário
# comentar o primeiro panic.
$ go run panic.go
panic: uhhh... Houston, we have a problem.

goroutine 1 [running]:
main.main()
	/.../panic.go:16 +0x27
exit status 2


# Note que, diferentemente de outras linguagens que 
# usam exceções para tratar erros, em Go é idiomático
# usar retornos indicativos de erro sempre que possível.