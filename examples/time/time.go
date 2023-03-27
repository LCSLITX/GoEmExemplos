// Go oferece suporte extenso para tempos e durações;
// Aqui estão alguns exemplos.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Aqui é capturado o horário atual.
	now := time.Now()
	p(now)

	// Pode ser construída uma struct `time` provendo
	// ano, mês, dia, etc. Tempos (ou horários) sempre
	// estão associados com uma localização,
	// por exemplo, fuso-horário.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// É possível extrair vários componentes do
	// valor horário.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Dia da semana, de Domingo a segunda, também está disponível
	// pela função `Weekday`.
	p(then.Weekday())

	// Estes métodos comparam dois tempos, testando se o primeiro
	// ocorre antes, depois ou ao mesmo segundo, respectivamente.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// O método `Sub` retorna a duração, `Duration`, representando
	// o intervalo entre dois tempos.
	diff := now.Sub(then)
	p(diff)

	// É possível computar o tamanho da duração em várias unidades.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Também se pode usar o `Add` para avançar o tempo por uma
	// duração determinada, ou com `-` para retroceder.
	p(then.Add(diff))
	p(then.Add(-diff))
}
