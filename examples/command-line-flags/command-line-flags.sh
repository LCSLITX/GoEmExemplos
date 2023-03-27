# Para testar o programa das flags, é melhor
# é melhor criar um binário com o comando `go build`
# antes.
$ go build command-line-flags.go

# Execute o programa passando valores para
# todas as flags.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Note que, se alguma flag for omitida, elas
# automaticamente são iniciadas com os valores
# padrão estipulados.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# argumentos podem ser fornecidos depois das flags.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Note que o pacote `flag` requer que todas as flags
# sejam usadas antes de argumentos. Do contrário, as
# flags serão interpretadas como argumentos.
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Usa-se as flags `-h` ou `--help` para acessar
# um texto de auxílio à utilização do programa.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Se for utilizada uma flag que não foi 
# especificada no pacote `flag`, o programa
# irá exibir uma mensagem de erro
# juntamente com um texto de auxílio.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
