package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "strings"
)

// handler greets the user
func handler(w http.ResponseWriter, r *http.Request) {
    name := strings.TrimPrefix(r.URL.Path, "/")
    if name == "" {
        name = "World"
    }
    fmt.Fprintf(w, "Hello, %s! Welcome to my GitHub page!", name)
}

// timeHandler displays the current server time
func timeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format(time.RFC1123)
    fmt.Fprintf(w, "Current server time: %s", currentTime)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/time", timeHandler)
    log.Println("Starting server on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
