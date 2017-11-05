package main

import (
    "log"
    "net/http"
    "os"
    "time"
    "fmt"
)

var restrictedRoute jwt.Auth

var restrictedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "secret")
})

var regularHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "hello")
})

func main() {
  authErr := jwt.New(&restrictedRoute, jwt.Options{
    SigningMethodString:   "RS256",
    PrivateKeyLocation:    "keys/app.rsa",     // `$ openssl genrsa -out app.rsa 2048`
    PublicKeyLocation:     "keys/app.rsa.pub", // `$ openssl rsa -in app.rsa -pubout > app.rsa.pub`
    RefreshTokenValidTime: 72 * time.Hour,
    AuthTokenValidTime:    15 * time.Minute,
    Debug:                 false,
    IsDevEnv:              true,
  })
  if authErr != nil {
    log.Println("Error initializing the JWT's!")
    log.Fatal(authErr)
  }

  http.HandleFunc("/", regularHandler)
  // this will never be available because we never issue tokens
  // see login_logout example for how to provide tokens
  http.Handle("/restricted", restrictedRoute.Handler(restrictedHandler))

  log.Println("Listening...")
    http.ListenAndServe(":"+os.Getenv("POST"), nil)
}
