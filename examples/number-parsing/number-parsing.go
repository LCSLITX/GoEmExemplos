// Parsear números a partir de strings é algo básico e
// muito comum em muitos softwares; aqui está como fazer em Go.

package main

// O pacote nativo `strconv` fornece algumas funções
// para parsear números.
import (
	"fmt"
	"strconv"
)

func main() {

	// Com `ParseFloat`, o segundo argumento informa à função
	// quanto bits de precisão serão utilizados ao parsear.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Para o `ParseInt`, o segundo argumento `0`, significa
	// em que base está o número a ser parseado. E o terceiro
	// argumento `64` requer que o resultado caiba em 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` reconhecerá números formatados em hexadecimal.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Também está disponível a `ParseUint` para parsear
	// inteiros unsigned, sem sinal.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` é uma função de conveniência para parsear
	// `int` em base-10.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// As funções retornam erro em caso de input errado.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
