# `São recebidos os valores `"um"` e então `"dois"`,
# conforme esperado.
$ time go run select.go 
recebido um
recebido dois
# Note que o tempo total de execução é aproximadamente 2
# segundos, pois ambas as funções `Sleep`, com 1 e
# 2 segundos executam concorrentemente.
real	0m2.245s
