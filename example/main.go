package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love golang!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server at port 8080...")
    http.ListenAndServe(":8080", nil)
}
