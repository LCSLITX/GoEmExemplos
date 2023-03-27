#  Ao executar `exit.go` usando `go run`, a saída
# será capturada pelo Go e então impressa.
$ go run exit.go
exit status 3

# Ao construir e executar um binário, é possível
# visualizar o status no terminal.
$ go build exit.go
$ ./exit
$ echo $?
3

# Note que o `!` do programa nunca foi exibido.
