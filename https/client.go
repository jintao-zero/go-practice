package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	pool := x509.NewCertPool()
	caPath := "ca.crt"
	caCrt, err := ioutil.ReadFile(caPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}
	tlsConfig := &tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{cliCrt},
	}
	tr := http.Transport{TLSClientConfig: tlsConfig}
	client := http.Client{Transport: &tr}
	url := "https://localhost/hello"
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}
