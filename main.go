package main

import (
    "net/http"
    "os"
)

var regularHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
})

func main() {
    http.HandleFunc("/", regularHandler)
    http.ListenAndServe(":"+os.Getenv("PORT") ,nil)
}
