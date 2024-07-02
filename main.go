package main

import (
    "encoding/json"
    "net/http"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
    headers := r.Header

    // Set Content-Type to application/json
    w.Header().Set("Content-Type", "application/json")

    // Convert headers to JSON
    jsonHeaders, err := json.Marshal(headers)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Write JSON response
    w.Write(jsonHeaders)
}

func main() {
    http.HandleFunc("/", headersHandler)
    http.ListenAndServe(":8080", nil)
}

