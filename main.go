package main

import (
    "net/http"
    "os"
    "fmt"
)

var regularHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "hello")
})

func main() {
  http.HandleFunc("/", regularHandler)
    
  http.ListenAndServe(":"+os.Getenv("POST"), nil)
}
