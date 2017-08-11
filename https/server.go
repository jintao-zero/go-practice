package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	pool := x509.NewCertPool()
	caPath := "ca.crt"
	caCrt, err := ioutil.ReadFile(caPath)
	if err != nil {
		fmt.Println("ReadFile err ", err)
		return
	}
	ok := pool.AppendCertsFromPEM(caCrt);
	if !ok {
		fmt.Println("append err")
		return
	}

	fmt.Println(ok)
	tlsConfig := &tls.Config{
		ClientCAs:pool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	server := http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}
	http.HandleFunc("/hello", HelloServer)
	err = server.ListenAndServeTLS("server.crt", "server.key")
	log.Fatal(err)
}
