// Go fornece suporte nativo para [base64
// encoding/decoding](https://en.wikipedia.org/wiki/Base64).

package main

// Esta sintaxe importa o pacote `encoding/base64` com o
// nome `b64` ao invés do nome padrão `base64`.
// Isso poupará algum espaço abaixo.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Aqui está a `string` que será objeto de encode/decode.
	data := "abc123!?$*&()'-=@~"

	// Go suporta tanto padrão quanto base64 compatível
	// com URL. Aqui está como realizar encode usando o
	// encoder padrão. O encoder requer `[]byte` então
	// antes é convertida a string para este tipo.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Realizar o Decode pode retornar um erro,
	// o qual pode ser checado se ainda não se sabe
	// se o input está bem formatado.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// Este encode/decode está usando base-64 compatível
	// com URL.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
