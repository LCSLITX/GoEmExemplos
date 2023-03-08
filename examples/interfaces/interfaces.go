// _Interfaces_ são coleções nomeadas de assinaturas
// de métodos.

package main

import (
	"fmt"
	"math"
)

// Aqui está uma interface básica para formas geométricas.
type geometria interface {
	area() float64
	perimetro() float64
}

// Para o exemplo, será implementado esta interface nos
// tipos `retangulo` e `circulo`.
type retangulo struct {
	largura, altura float64
}
type circulo struct {
	raio float64
}

// Para implementar uma interface em Go, apenas é necessário
// implementar a uma struct todos os métodos constantes numa
// interface. Aqui é implementada a interface `geometria` em `retangulo`.
func (r retangulo) area() float64 {
	return r.largura * r.altura
}
func (r retangulo) perimetro() float64 {
	return 2*r.largura + 2*r.altura
}

// E aqui, a implementação em `circulo`.
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// Se uma variável tem um tipo interface, então é possível
// chamar os métodos constantes daquela interface.
// Aqui uma função genérica `medir`, usando este conceito
// para funcionar com qualquer struct que implemente a
// interface `geometria`.
func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	r := retangulo{largura: 3, altura: 4}
	c := circulo{raio: 5}

	// Ambas as structs `circulo` e `retangulo` implementam
	// a interface `geometria`, então é possível utilizar
	// instâncias destas structs como argumentos para a função
	// `medir`.
	medir(r)
	medir(c)
}
