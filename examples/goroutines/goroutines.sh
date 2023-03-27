# Ao rodar este programa, primeiro é visualizada a saída
# das chamadas bloqueantes, síncronas, e só então a saída 
# das duas goroutines. A saída de goroutines podem ser
# exibidas em ordem diversa da disposta no código,
# justamente porque goroutines são executadas 
# concorrentemente pelo runtime de Go.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# Em seguida veremos um complemento às goroutines
# em programas Go que utilizam concorrência: canais.
