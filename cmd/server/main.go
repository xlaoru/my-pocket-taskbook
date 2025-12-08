package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func main() {
	/* storage, err := db.New()

	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	defer storage.Pool.Close()

	fmt.Println("Connected to PostgreSQL!")

	if err := storage.Migrate(); err != nil {
		log.Fatalf("DB migration failed: %v", err)
	} */

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is working!")
	})

	fmt.Println("Server listening to", PORT)
	http.ListenAndServe(PORT, mux)
}
