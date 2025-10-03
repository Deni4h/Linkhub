package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Ambil PORT dari environment, default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	http.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "testing"
		}
		fmt.Fprintf(w, "ini : , %s!", name)
	})

	http.HandleFunc("/health/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("alive"))
	})

	http.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		// di real app bisa cek koneksi DB dsb
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ready"))
	})

	log.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
