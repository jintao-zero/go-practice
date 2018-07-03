package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/urfave/negroni"
)

func index(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome!")
}
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", index)
    n := negroni.Classic()
    n.UseHandler(router)
    n.Run(":8080")
}
