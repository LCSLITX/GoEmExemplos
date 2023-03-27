// URLs possibilitam [uma forma uniforme de encontar recursos](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Aqui está como parsear URLs em Go.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// Aqui será parseado esta URL de exemplo, que inclui
	// schema, informação de autênticação, host, porta,
	// caminho (path), query params, e fragmentos de query.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Aqui a URL é parseada e assegurada de que não
	// ocorreram erros.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Para acessar o schema é bem direto.
	fmt.Println(u.Scheme)

	// `User` contem toda informação de autenticação;
	// `Username` e `Password` exibem a respectiva valor.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// O `Host` contem tanto hostname quanto a porta,
	// se presente. A função `SplitHostPort` separa
	// os dados.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Aqui são extraídos o `path` e o fragmento
	// depois de `#`.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Para obter os query params em uma string de formato
	// `chave=valor`, usa-se `RawQuery`. Também é possível
	// parsear query params para um `map`.
	// Os maps de query params parseados são formatados em
	// `map[string][]string` então, para acessar o primeiro
	// valor da primeira chave, é necessário utilizar `[0]`.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"])
	fmt.Println(m["k"][0])
}
