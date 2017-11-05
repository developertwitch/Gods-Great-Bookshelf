package main

import (
    "net/http"
    "log"
    "time"
    "os"
    "fmt"

    "github.com/adam-hanna/jwt-auth/jwt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello")
}

func main() {
    http.HandleFUnc("/", handler)
    httpe.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
