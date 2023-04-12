# Go Em Exemplos

Este repositório contém conteúdos e ferramentas de build para o site [Go Em Exemplos](https://lcslitx.github.io/GoEmExemplos/),
um site que ensina Go com exemplos comentados.

### Vista Geral

O site Go Em Exemplos é construído extraindo o código e os
comentários dos arquivos no diretório `examples` e, utilizando
como base os `templates`, cria uma versão estática no
diretório `public`.
Os programas utilizados para este processo de construção podem
ser encontrados no diretório `tools`, juntamente com as dependências
especificadas no arquivo `go.mod`.

O diretório `public` pode ser servido por qualquer sistema
de conteúdo estático. 
O site de produção do [Gobyexample.com](gobyexample.com),
por exemplo, utiliza S3 e CloudFront. A [versão em Português](https://lcslitx.github.io/GoEmExemplos/)
Brasileiro é servida pelo [Github Pages](https://pages.github.com/).

### Building

Para realizar o build do site, é necessário ter Go instalado. Execute:

```console
$ tools/build
```

Ou para realizar o build continuamente em um loop:

```console
$ tools/build-loop
```

Para hospedar o site localmente:

```
$ tools/serve
```

É possível acessar pela porta 8000 (`http://127.0.0.1:8000/`) no navegador.

### Publicação

Para fazer upload do site utilizando AWS:

```console
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ tools/upload
```

### Licenças

O copyright pertence ao Mark McGranaghan e seu uso está licenciado sob a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

O copyright do mascote Gopher pertence ao [Renée French](https://reneefrench.blogspot.com/) e está licenciado sob a 
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Traduções

Traduções dos contribuidores do GoByExample estão disponíveis em:

* [Chinês](https://gobyexample-cn.github.io/) por [gobyexample-cn](https://github.com/gobyexample-cn)
* [Tcheco](http://gobyexamples.sweb.cz/) por [martinkunc](https://github.com/martinkunc/gobyexample-cz)
* [Francês](http://le-go-par-l-exemple.keiruaprod.fr) by [keirua](https://github.com/keirua/gobyexample)
* [Italiano](https://gobyexample.it) pela [comunidade Go Italian](https://github.com/golangit/gobyexample-it)
* [Japonês](http://spinute.org/go-by-example) por [spinute](https://github.com/spinute)
* [Coreano](https://mingrammer.com/gobyexample/) por [mingrammer](https://github.com/mingrammer)
* [Russo](https://gobyexample.com.ru/) por [badkaktus](https://github.com/badkaktus)
* [Espanhol](http://goconejemplos.com) pela [comunidade Go Mexico](https://github.com/dabit/gobyexample)
* [Ucraniano](https://butuzov.github.io/gobyexample/) por [butuzov](https://github.com/butuzov/gobyexample)
* [Português Brasileiro](https://lcslitx.github.io/GoEmExemplos/) por [lcslitx](https://github.com/LCSLITX)

### Thanks

Agradecimento ao [Jeremy Ashkenas](https://github.com/jashkenas)
pelo [Docco](http://jashkenas.github.io/docco/), que inspirou este projeto.

### FAQ

#### Encontrei um problema com os exemplos; o que eu faço?

Ficamos felizes em corrigir problemas e aceitar contribuições! Por favor, submeta um
[issue](https://github.com/LCSLITX/GoEmExemplos/issues) ou envie uma Pull Request.
Leia o arquivo `CONTRIBUINDO.md` para mais detalhes.

#### Qual versão de Go é necessária para executar os exemplos?

Graças à [retrocompatibilidade](https://go.dev/doc/go1compat) da linguagem Go,
é esperado que a grande maioria dos exemplos funcinoem nas versões mais recentes de Go, bem como nas versões menos atualizadas.

Considerando isto, alguns exemplos demonstram recursos recém implementados; portanto,
é recomendado executar com as versões oficiais mais recentes.
(Veja o [histórico de releases](https://go.dev/doc/devel/release) para mais detalhes).

#### Quando executo o exemplo em minha máquina, a ordem do output é diferente. Algo está errado?

Alguns dos exemplos demonstram execução de código concorrente, o qual tem uma ordem de execução não determinística. Tudo depende de como a Execução (runtime) do Go organiza as goroutines. Isto pode depender do Sistema operacional, arquitetura do CPU e até de versão de Go.

De maneira similar, exemplos que iteram sobre maps podem produzir itens em ordens diferentes da que se está obtendo na sua máquina. Isto acontece porque a ordem de iteração nos maps de Go [não são específicas e não tem garantia de serem iguais às iterações anteriores ou posteriores](https://go.dev/ref/spec#RangeClause).

Isto não significa que há algo de errado com o exemplo. Tipicamente, o código neste exemplos não são sensíveis à ordem do output; se o código é sensível à essa ordem, provavelmente trata-se de um bug, então sinta-se livre para reportar.
