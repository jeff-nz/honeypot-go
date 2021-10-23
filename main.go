package main

import (
    "log"
    "net/http"
)


func main() {
    log.Print("Starting up on honeypot-go")

    http.HandleFunc("/", handleRandomGetResponse)

    log.Print("Listening on localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
