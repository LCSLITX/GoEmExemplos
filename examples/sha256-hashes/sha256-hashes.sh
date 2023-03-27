# Ao executar o código, é exibida a hash em formato hex
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# Também é possível computar outras hashes usando padrão 
# similar visto anteriormente. Por exemplo, para computar
# hashes SHA512 é preciso importar `crypto/sha512` e
# utilizar `sha512.New()`.

# Note que se você precisa de hashes criptograficamente 
# seguras é bom pesquisar um pouco sobre
# [hash strength](https://en.wikipedia.org/wiki/Cryptographic_hash_function)!
