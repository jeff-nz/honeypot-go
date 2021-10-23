package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)


func renderHoneypotResponse(w http.ResponseWriter, r *http.Request) {
    log.Printf("renderHoneypotResponse() called")
    flusher, ok := w.(http.Flusher)
    if !ok {
        panic("expected http.ResponseWriter to be an http.Flusher")
        http.Error(w, "Method not allowed", http.StatusInternalServerError)
        return
    }

    for true {
        log.Printf("has rendered dummy content")
        fmt.Fprintln(w, "NOT FOUND")
        flusher.Flush()
        time.Sleep(5 * time.Second)
    }
}
