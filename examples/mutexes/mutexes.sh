# Executando o programa, é exibido que os
# contadores foram atualizados, conforme esperado.
$ go run mutexes.go
map[a:20000 b:10000]

# Em seguida, será apresentado como implementar este
# mesmo gerenciamento de estado de tarefas usando
# apenas goroutines e canais.