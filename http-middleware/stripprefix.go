package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("root", r.RequestURI)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test", r.RequestURI)
}

func main() {
	fs := http.FileServer(http.Dir("/tmp"))
	//http.Handle("/", (http.HandlerFunc)(root))
	//http.Handle("/test/", (http.HandlerFunc)(test))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.ListenAndServe(":8080", nil)
}
