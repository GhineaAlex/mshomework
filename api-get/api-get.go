package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler function for the root path
func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Just a GET", http.StatusMethodNotAllowed)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "This was a GET")
}

func main() {
    http.HandleFunc("/", handler)

    fmt.Println("Starting server")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
