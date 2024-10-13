package main

import (
    "fmt"
    "net/http"
	"encoding/json"
)

type Response struct {
    Message string `json:"message"`
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Allow all origins for this example; adjust as needed
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        // Handle preflight requests
        if r.Method == http.MethodOptions {
            return
        }

        next.ServeHTTP(w, r) // Call the next handler
    })
}

func handler(w http.ResponseWriter, r *http.Request) {
    // Set the response header to application/json
    w.Header().Set("Content-Type", "application/json")

	fmt.Println("New ping request")

    response := Response{
        Message: "Hello, World!",
    }

    json.NewEncoder(w).Encode(response)
}

func main() {
	var mux = http.NewServeMux()
    mux.HandleFunc("/ping", handler)
    fmt.Println("Starting server on :8080...")

    if err := http.ListenAndServe(":8080", corsMiddleware(mux)); err != nil {
        fmt.Println(err)
    }
}
