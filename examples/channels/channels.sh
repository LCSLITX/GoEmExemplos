# Ao rodar o programa a mensagem `"ping"` é passada com 
# successo de uma goroutine para outra via o canal que
# foi criado.
$ go run channels.go 
ping

# Por padrão, ambos os canais de envio e recebimento 
# ficam bloqueados até que ambos estejam prontos. 
# Esta propriedade permite aguardar ao final do programa 
# pela mensagem sem que seja preciso qualquer outro tipo
# de sincronização.
