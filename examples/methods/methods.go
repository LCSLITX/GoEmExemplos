// Go suporta _métodos_ definidos em structs.

package main

import "fmt"

type rect struct {
	width, height int
}

// O método `area` tem um _receiver type_ de `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Métodos podem ser definidos tanto com _receiver types_ que são
// ponteiros ou valores. Aqui está um exemplo de receiver de valor.
func (r rect) perimetro() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Aqui são chamados dois métodos definidos para a struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perimetro:", r.perimetro())

	// Go faz automaticamente conversões entre valores e
	// ponteiros para chamadas de métodos. Geralmente se
	// usa um _pointer receiver_ para se evitar a necessidade
	// de passar uma cópia da struct como parâmetro da função
	// ou para permitir que o método altere a struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimetro:", rp.perimetro())
}
