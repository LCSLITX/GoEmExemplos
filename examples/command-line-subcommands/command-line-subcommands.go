// Alguns programas de linha de comando, como o próprio `go` ou `git`
// possuem muitos *subcomandos*, cada um com seu próprio conjunto
// de flags. Por exemplo, `go build` e `go get` são dois subcomandos
// diferentes de `go`.
// O pacote `flag` permite definir subcomandos com seu próprio
// conjunto de flags.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Um subcomando é declarado usando a função `NewFlagSet`
	// e, em seguida, se prossegue com a definição das
	// flags específicas para este subcomando.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Para diferentes subcomandos podem ser definidas
	// diferentes flags suportadas.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// O subcomando é esperado como primeiro argumento
	// do programa.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Verifica qual subcomando é utilizado.
	switch os.Args[1] {

	// Para cada subcomando, são parseadas suas próprias flags e
	// tem acesso aos eventuais argumentos adicionais ao final.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
