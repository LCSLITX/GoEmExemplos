# É esperado que o código realize exatamente
# 50.000 operações.
# Se fosse utilizada a forma não-atômica de
# incrementação, provavelmente o resultado
# seria diferente entre as execuções, porque
# as goroutines interfeririam umas com as
# outras. Alẽm disso, provavelmente aconteceria
# falhas de data race, que é possível visualizar
# executando o código com a flag `-race` 
# (`go run -race atomic-counters.go`)
$ go run atomic-counters.go
ops: 50000


# Código com `ops++` ao invés de 
# `atomic.AddUint64(&ops, 1)`
# sendo executado com a flag `-race`
$ go run -race atomic-counters.go
==================
WARNING: DATA RACE
Read at 0x00c00001a0f8 by goroutine 10:
  main.main.func1()
      /.../atomic-counters.go:38 +0x46
==================
ops: 6573
Found 1 data race(s)
exit status 66


# Em seguida, será apresentado mutexes, 
# outra ferramenta para gerenciar estado.