// Algumas vezes é necessário ordenar uma coleção por algo
// que não sua ordem natural. Por exemplo, suponha que
// queiramos ordenar uma string pelo seu tamalho ao invés
// de alfabeticamente. Aqui está um exemplo de ordenação
// personalizada em Go.

package main

import (
	"fmt"
	"sort"
)

// Para ordenar utilizando uma função própria em Go, é
// necessário um tipo correspondente. Aqui é criado um
// tipo `byLength` que é apenas um apelido (alias) para
// o tipo nativo `[]string`.
type byLength []string

// Aqui, foi feita a implementação da Interface específica
// do pacote sort `sort.Interface`, com os métodos `Len`,
// `Less`, e `Swap` - para ser possível utilizar a função
// genérica `Sort` do pacote. `Len` e `Swap`
// são geralmente similares entre os tipos e `Less` será
// responsável pela lógica de ordenação. Neste caso, as
// strings serão ordenadas de forma crescente, pelo seu
// tamanho, então são utilizados `len(s[i])` e `len(s[j])`.
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// Com tudo isso pronto, já podemos utilizar nossa ordenação
// personalizada. Para isso, apenas convertemos o slice
// original `fruits` para um do tipo `byLength` que acabou de
// ser criado, e então usar o `sort.Sort` neste novo slice.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
