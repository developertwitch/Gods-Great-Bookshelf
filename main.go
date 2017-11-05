package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Your port: "+os.Getenv("PORT"))
}
