package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

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
    fmt.Fprintf(responseWriter, "Hello astaxie!")
}

func main() {
    http.HandleFunc("/", sayhelloName)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}