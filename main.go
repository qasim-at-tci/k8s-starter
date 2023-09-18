package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
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
