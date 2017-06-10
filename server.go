package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

const server_port = 9090

func handleDefault(responseWriter http.ResponseWriter, request *http.Request) {
    request.ParseForm()
    fmt.Println(request.Form)
    fmt.Println("path", request.URL.Path)
    fmt.Println("scheme", request.URL.Scheme)
    fmt.Println(request.Form["url_long"])
    for k, v := range request.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(responseWriter, "Hello World!")
}

func main() {
    fmt.Println("Starting server on port:", server_port)
    http.HandleFunc("/", handleDefault)
    err := http.ListenAndServe(fmt.Sprintf(":%d", server_port), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}