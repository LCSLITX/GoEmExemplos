$ go run channel-synchronization.go      
working...done                  

# Ao remover o recebimento `<- done` do código, o
# programa será finalizado antes que a _worker goroutine_
# comece a ser executada.
