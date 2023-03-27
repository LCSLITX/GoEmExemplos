# O programa executado exibe 5 tarefas sendo executadas
# por vários workers. O programa leva apenas 2 segundos
# para finalizar, ainda que as tarefas totalizem 5
# segundos, por ter 3 workers trabalhando 
# concorrentemente.
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
