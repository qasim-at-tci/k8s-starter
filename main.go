package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Status struct {
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/_/healthz", healthCheckHandler)
	log.Println("Starting server - Press CTRL+C to exit..")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("ERROR!")
	}
	fmt.Fprintf(w, "Hello from %s", hostname)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	status := Status{Status: "OK"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
