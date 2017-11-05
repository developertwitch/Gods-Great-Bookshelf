package main

import (
    "net/http"
    "os"
    "fmt"
)

func handler(w http.Response Writer, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
}

func main() {
  http.HandleFunc("/", regularHandler)
    
  http.ListenAndServe(":"+os.Getenv("POST"), nil)
}
