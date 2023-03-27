// Algo muito comum em diversos softwares é a necessidade
// de saber o número de segundos, milisegundos ou nanosegundos
// desde a [era Unix](https://pt.wikipedia.org/wiki/Era_Unix).
// Aqui está como isso pode ser feito em Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// O `time.Now` com `Unix`, `UnixMilli` ou `UnixNano`
	// retornam o tempo transcorrido desde a Era Unix,
	// em segundos, milisegundos ou nanosegundos,
	// respectivamente.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// Também é possível converter segundos ou nanosegundos
	// de inteiros para o `time` ou tempo correspondente.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
