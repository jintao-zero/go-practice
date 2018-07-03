package main

import (
    "github.com/urfave/negroni"
    "net/http"
    "fmt"
)

func index(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Welcome!")
}

func main()  {
    mux := http.NewServeMux()
    mux.HandleFunc("/", index)
    n := negroni.Classic()
    n.UseHandler(mux)
    n.Run(":8080")
}
