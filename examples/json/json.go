// Go oferece suporte nativo para encoding e decoding
// de JSON, inclusive _de_ e _para_ tipos de dados nativos e
// personalizados.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Serão utilizadas estas duas structs para demonstrar
// encoding e decoding de tipos personalizados abaixo.
type response1 struct {
	Page   int
	Fruits []string
}

// Apenas campos exportados serão encoded/decoded em JSON.
// Campos devem começar com letras maiusculas para serem exportados.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Primeiro, será apresentado encoding básico de
	// tipos de dados para strings JSON. Aqui estão
	// alguns exemplos para valores primitivos.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// E aqui estão alguns para slices e maps,
	// que, após o encode, correspondem a arrays
	// e objetos JSON, como esperado.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// O pacote JSON pode realizar encode de tipos de dados
	// personalizados de maneira automática. Serão incluídos
	// na saída, apenas os campos exportados e seus nomes
	// serão utilizados como chaves no objeto JSON por padrão.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// É possível utilizar tags nas declarações dos campos
	// da struct para customizar os nomes das chaves na saída
	// JSON. Veja a definição de `response2` acima para entender
	// o exemplo das tags.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Agora, será apresentado o decoding de dado em
	// JSON para valores em Go. Aqui está um exemplo
	// para uma estrutura genérica.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// É necessário fornecer a variável para onde o pacote
	// JSON deve inserir os dados decoded. Este
	// `map[string]interface{}` armazenará um map de strings
	// para tipos arbitrários de dados.
	var dat map[string]interface{}

	// Aqui é que acontece o decoding, de fato. E uma
	// verificação de eventuais erros associados ao
	// decoding.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Para usar os valores no map, é necessário
	// convertê-los para o tipo apropriado.
	// Por exemplo, aqui é convertido o valor em `num`
	// para o tipo esperado `float64`.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Acessar dados aninhados requer uma serie de
	// conversões.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Também é possível usar decode de JSON em tipos de
	// dados customizados. Isto adiciona a vantagem de
	// ter maior seguraça dos tipos, ou `type-safety`
	// no código, eliminando a necessidade de asserções
	// ao acessar dados decoded.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// Nos exemplos acima sempre foram utilizados bytes e
	// strings como intermediários entre os dados e a
	// representação JSON nas saídas. Também é possível
	// utilizar encoding JSON em stream diretamente no
	// `os.Writer`, como o `os.Stdout` ou até no corpo
	// de respostas HTTP.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
