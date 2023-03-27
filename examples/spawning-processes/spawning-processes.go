// Por vezes é necessário que programas em
// Go invoquem processos ou programas _não-Go_.

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// De início será utilizado um simples comando que não recebe
	// nenhum argumento ou input e apenas imprime algo no `stdout`.
	// O comando `exec.Command` cria um objeto para representar
	// este processo externo.
	dateCmd := exec.Command("date")

	// O método `Output` executa o comando, aguarda sua finalização
	// e coleta o output padrão.
	// Se não houver nenhum erro, `dateOut` terá os bytes com informação
	// da data.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` e outros métodos de `Command` retornarão
	// um `*exec.Error` se houver algum problema ao executar
	// o comando, como por exemplo caminho errado, e
	// `*exec.ExitError` se o comando foi executado mas
	// retornou com um código não-zero.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// Agora será apresentado um caso um pouco mais complexo
	// em que é enviado um dado para o processo externo em seu
	// `stdin`, através de um pipe e coletado o resultado a
	// partir de seu `stdout` também a partir de um pipe.
	grepCmd := exec.Command("grep", "hello")

	// Aqui se instancia os pipes, inicializa o próprio
	// processo, então escreve e envia algum input, lê o
	// resultado de saída e, finalmente, aguarda o processo
	// finalizar.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// Verificação de erros foram omitidas no exemplo acima,
	// mas é possível utilizar o padrão usual `if err != nil`
	// para todos eles. Também foi coletado apenas o resultado
	// de `StdoutPipe`, mas é possível também coletar de
	// `StderrPipe` exatamente do mesmo jeito.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Note que, ao invocar comandos, é necessário fornecer
	// um comando explicitamente delineado e um array de argumentos
	// ao invés de passar apenas uma string de linha de comando
	// command-line string. Se a intenção for gerar um comando
	// com uma string, é possível utilizar a flag `-c` do `bash`.
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
