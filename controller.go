package main

import (
    "log"
    "net/http"
)

func handleRandomGetResponse(w http.ResponseWriter, r *http.Request) {
    log.Printf("handleRandomGetResponse() called")
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    renderHoneypotResponse(w, r)
}
