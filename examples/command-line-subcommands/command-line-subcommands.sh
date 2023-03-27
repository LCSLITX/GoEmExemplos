$ go build command-line-subcommands.go 

# Utilização do primeiro subcomando
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# Agora o segundo.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# O segundo com tentativa de utilizar flags do primeiro.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# Em seguida serão apresentadas as variáveis de ambiente,
# outra forma muito comum de parametrizar programas.
