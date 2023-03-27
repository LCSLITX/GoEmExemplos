// Começando na versão 1.18, Go adicionou suporte a
// _generics_, também conhecidos como _type parameters_
// ou _tipos de parâmetros_.

package main

import "fmt"

// Como um exemplo de função genérica, `MapKeys`
// recebe um map de qualquer tipo e retorna um slice
// de suas chaves.
// Essa função tem dois tipos de parâmetros, `K` e `V`;
// K possui a _constraint_, ou _restrição_, `comparable`/
// `comparável`, ou seja, é possível comparar valores deste tipo
// com os operadores `==` e `!=`. Isto é requerido para chaves
// de maps em Go.
// `V` possui a _constraint_ `any`/
// `qualquer`, ou seja, não é restringida a nenhum tipo específico
// (`any` é um alias para `interface{})`.
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Como um exemplo de tipo genérico, `List` é uma
// `singly-linked list` com valores de qualquer tipo.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// É possível definir métodos para tipos genéricos
// assim como é feito em tipos regulares, mas é
// preciso manter o tipo do parâmetro.
// O tipo é `List[T]`, não `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// Ao invocar funções genéricas, é possível aproveitar
	// _type inference_ ou _inferência de tipo_. Note que
	// não é preciso especificar os tipos para `K` e `V`
	// ao chamar `MapKeys` - o compilador infere seus
	// tipos automaticamente.
	fmt.Println("keys:", MapKeys(m))

	// ... embora também seja possível especificá-los
	// explicitamente.
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
