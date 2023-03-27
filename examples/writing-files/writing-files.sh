# Tente executar o código de escrita em arquivo.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# E então, verifique o conteúdo dos arquivos escritos.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# Em seguida, serão apresentadas algumas aplicações
# de entrada e saída de arquivos que foram vistas
# anteriormente, em streams de `stdin` e `stdout`.
