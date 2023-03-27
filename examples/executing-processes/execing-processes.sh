# Ao rodar o programa, ele é substituído pelo `ls`.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Note que Go não oferece a função `fork`,
# clássica do Unix. Usualmente isto não é um
# problema, porque goroutines, invocação e
# executação de processos cobrem a maioria dos
# casos de uso de `fork`.