package main

import (
    "net/http"
    "os"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
