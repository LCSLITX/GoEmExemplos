// Go oferece suporte nativo para formato XML
// e formatos semelhates com o pacote `encoding.xml`.

package main

import (
	"encoding/xml"
	"fmt"
)

// A struct Plant será mapeada para XML. Similar aos
// exemplos JSON, as tags dos campos possuem diretivas
// para o encoder e decoder. Aqui são usados alguns
// recursos especiais do pacote XML: O campo `XMLName`
// dita o nome do elemento XML representando a struct;
// `id,attr` significa que o campo `Id` é um _attributo_
// XML ao invés de um elemento aninhado.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Gera um XML representando a struct plant;
	// ao usar `MarshalIndent`, é produzida uma saída
	// mais legível para humanos.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Para adicionar um header XML genérico à saída,
	// é necessário acrescentar explícitamente.
	fmt.Println(xml.Header + string(out))

	// O `Unmarshal` pode ser usado para parsear um stream
	// de bytes com XML em uma estrutura de dados. Se o XML
	// estiver mal formado ou não puder ser mapeado na estrutura
	// plant, um erro descritivo será retornado.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// O tag `parent>child>plant` do campo Plants, diz ao encoder
	// para aninhar todas as plantas sob `<parent><child>...`
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
