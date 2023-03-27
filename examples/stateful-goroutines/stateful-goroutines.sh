# Ao executar este código, é exibido que o exemplo de
# gerenciamento de estado baseado em goroutine completa
# cerca de 80.000 operações no total.
$ go run stateful-goroutines.go
readOps: 71708
writeOps: 7177

# Para este caso em particular, o exemplo baseado em 
# goroutines é um pouco mais acoplado que o baseado em
# mutex. Embora possa ser útil em alguns casos, como 
# por exemplo onde exista outros canais envolvidos ou 
# ao gerenciar múltiplos mutex, que seria mais propenso 
# a erros. O correto é utilizar a forma que for mais 
# natural, especialmente no que diz respeito à comprensão
# da forma que faça mais sentido para a realidade do
# código.

