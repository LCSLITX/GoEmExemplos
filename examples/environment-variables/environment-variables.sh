# Executando o programa, é exibido  que o valor de `FOO`
# é o definido no código, mas BAR está vazio.
$ go run environment-variables.go
FOO: 1
BAR: 

# A lista de chaves no ambiente vai variar
# a depender da sua máquina.
TERM_PROGRAM
PATH
SHELL
...
FOO

# É possível definir o valor para uma variável ao executar
# o programa, assim `BAR` terá um valor.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
