package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	config, err := os.ReadFile("/etc/example.cfg")
	if err != nil {
		log.Fatal(err)
	}

	// Define handler that returns "Hello $ENV"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello "))
		w.Write(config)
	})

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
