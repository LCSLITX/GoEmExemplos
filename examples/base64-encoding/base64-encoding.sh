# As strings resultantes de ambos os métodos de encode,
# padrão e compatível com URL, são um pouco diferentes
# entre si (variando `+` e `-` ao final) mas se for
# realizado decode, ambas retornam a mesma string
# original, como esperado.
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
