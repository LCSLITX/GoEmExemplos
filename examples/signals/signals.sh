# Ao executar este código, ficará bloqueado aguardando
# um sinal. Ao digitar `ctrl-C` (que o terminal exibe
# como `^C`) é enviado um sinal `SIGINT`, que dispara
# a impressão da palavra `interrupt` e a
# interrupção do programa.
$ go run signals.go
awaiting signal
^C
interrupt
exiting
