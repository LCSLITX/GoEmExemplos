# Ao executar o código é exibida uma lista ordenada
# de strings pelo seu tamanho, conforme esperado.
$ go run sorting-by-functions.go 
[kiwi peach banana]

# Seguindo este mesmo padrão de criação de um tipo 
# personalizado, implementando os três métodos da
# interface daquele tipo, e então chamando sort.Sort
# numa coleção do respetivo tipo personalizado,
# é possível ordenar slices com funções arbitrárias.
