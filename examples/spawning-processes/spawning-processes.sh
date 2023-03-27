# Os processos invocados retornam o mesmo output que  
# retornaria se fosse executado diretamente da
# linha de comando.
$ go run spawning-processes.go 
> date
Thu 05 May 2022 10:10:12 PM PDT

# `date` não possui uma flag `-x`, então será finalizado
# com mensagem de erro e código de retorno não-zero.
command exited with rc = 1
> grep hello
hello grep

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go
