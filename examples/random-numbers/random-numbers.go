// O pacote `math/rand` fornece um gerador de
// [números pseudoaleatórios](https://pt.wikipedia.org/wiki/Gerador_de_n%C3%BAmeros_pseudoaleat%C3%B3rios).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Por exemplo, `rand.Intn` retorna um `int` `n` aleatório,
	// em que `0 <= n < 100`.
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `rand.Float64` retorna um `float64` `f`,
	// em que `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// Isto pode ser usado para gerar floats aleatórios em
	// outros intervalos, por exemplo `5.0 <= f < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// O gerador de números padrão é determinístico,
	// isso quer dizer que sempre irá produzir a mesma
	// sequência de números a cada vez, por padrão.
	// Para produzir números variados, é importante utilizar
	// uma `seed` que também varie. Note que NÃO é seguro
	// utilizar este método para números aleatórios que devam
	// ser secretos. Se esta for a intenção, recomenda-se a
	// utilização de `crypto/rand`.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Chama o `rand.Rand` resultante da criação anterior,
	// da nova instância com seed variante, exatamente como
	// a função `rand` do pacote.
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// Se uma seed constante é fornecida, a aleatoriedade
	// produz sempre a mesma sequência de números.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
