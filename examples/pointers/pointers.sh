# `zeroval` não altera o valor de `i` em `main`, 
# mas `zeroptr` altera porque recebe uma referência
# ao endereço de memória para aquela variável.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
