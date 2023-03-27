# Para executar o filtro de linha, primeiro
# crie um arquivo de texto com linhas em
# letras minusculas.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# Então use o filtro de linha para ter as
# linhas em letras maiusculas.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
