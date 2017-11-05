package main

import (
    "net/http"
    "os"
)

var restrictedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Welcome to the secret area!"))
})

var regularHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
})

func main() {
    http.HandleFunc("/", regularHandler)
    http.Handle("/restricted", restrictedHandler)

    
    http.ListenAndServe(":"+os.Getenv("PORT") ,nil)
}
