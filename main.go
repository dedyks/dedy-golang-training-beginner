package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", helloWorld)
	http.HandleFunc("/health", healthCheck)

	port := ":5050"
	log.Printf("Server started on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world"}`))
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"healthy"}`))
}
