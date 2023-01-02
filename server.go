package main

import (
    "fmt"
    "net/http"
    "example/greetings"
)

func main() {
    const PORT = ":1234"
    
    http.HandleFunc("/", func(w  http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "welcome to my website ")
    })
    
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(PORT, nil)
    
    fmt.Println(greetings.Hello("arham"),"Server running on Port", PORT)
}