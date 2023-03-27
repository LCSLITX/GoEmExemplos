package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	port := "8000"
	publicDir := "public"
	fmt.Printf("Servindo Go Em Exemplos em http://127.0.0.1:%s\n", port)

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
			fmt.Println(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	fmt.Printf("Acessível também pela rede em http://%s:%s\n", localAddr.IP, port)
	http.ListenAndServe(":"+port, http.FileServer(http.Dir(publicDir)))
}
