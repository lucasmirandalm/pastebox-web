package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasmirandalm/pastebox-web/internal/config"
)

func main() {
	cfg := config.Load()

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Pastebox is running")
	})

	addr := ":" + cfg.Port

	fmt.Printf("server running on http://localhost%s\n", addr)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
