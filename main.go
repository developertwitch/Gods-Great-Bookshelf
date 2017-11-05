package main

import (
    "net/http"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":6000", nil)
}
