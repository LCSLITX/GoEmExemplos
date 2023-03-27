# Executando o código de parseamento de URLs, são exibidas
# todas as partes extraídas.
$ go run url-parsing.go 
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
chave=valor&outrachave=outrovalor
map[chave:[valor] outrachave:[outrovalor]]
[valor]
valor