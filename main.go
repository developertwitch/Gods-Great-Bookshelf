package main

import (
    "log"
    "net/http"
    "os"
    "time"
    
    "github.com/adam-hanna/jwt-auth/jwt"
)

var restrictedRoute jwt.Auth

var restrictedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Welcome to the secret area!"))
})

var regularHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
})

func main() {
    authErr := jwt.New(&restrictedRoute, jwt.Options {
        SigningMethodStrin: "RS256",
        PrivateKeyLocatioPrivateKeyLocation: "keys/app.rsa",     // `$ openssl genrsa -out app.rsa 2048`
        PublicKeyLocatioPublicKeyLocation: "keys/app.rsa.pub", // `$ openssl rsa -in app.rsa -pubout > app.rsa.pub`
        RefreshTokenValidTimRefreshTokenValidTime: 72 * time.Hour,
        AuthTokenValidTimAuthTokenValidTime: 15 * time.Minute,
        DebuDebug: false,
        IsDevEnIsDevEnv: true,
    })
    if authErr != nil {
        log.Println("Error initializing the JWT's!")
        log.Fatal(authErr)
    }
    
    http.HandleFunc("/", regularHandler)
    http.Handle("/restricted", restrictedHandler)

    
    http.ListenAndServe(":"+os.Getenv("PORT") ,nil)
}
