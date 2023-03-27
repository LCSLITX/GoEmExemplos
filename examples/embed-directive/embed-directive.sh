# Utilize estes comandos para executar o exemplo.
# (Note: devida à limitação no playground Go,
# este exemplo apenas pode ser executado na sua maquina.)
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456

