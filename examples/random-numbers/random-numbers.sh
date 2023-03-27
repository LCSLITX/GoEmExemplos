# A depender de onde esta amostra for executada, alguns 
# números gerados podem ser diferentes. É importante notar
# que no [playground de Go](https://go.dev/play/p/C99PXOcv7G0), utilizar seed com `time.Now()`
# ainda produz resultados determinísticos, pela forma que
# o playground é implementado.
$ go run random-numbers.go
81,87
0.6645600532184904
7.123187485356329,8.434115364335547
0,28
5,87
5,87


# Veja a documentação do pacote
# [`math/rand`](https://pkg.go.dev/math/rand)
# para mais referências sobre outros
# recursos randômicos que Go pode fornecer.
