package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/lucasmirandalm/pastebox-web/internal/config"
	"github.com/lucasmirandalm/pastebox-web/internal/database"
	"github.com/lucasmirandalm/pastebox-web/internal/paste"
	"github.com/lucasmirandalm/pastebox-web/internal/render"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Open(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	renderer, err := render.New("ui/templates")
	if err != nil {
		log.Fatalf("failed to load templates: %v", err)
	}

	pasteHandler := paste.NewPasteHandler(renderer)

	r := chi.NewRouter()

	r.Get("/", pasteHandler.Home)

	r.Get("/health/db", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			http.Error(w, "database unavailable", http.StatusServiceUnavailable)
			return
		}

		fmt.Fprintln(w, "database OK")
	})

	addr := ":" + cfg.Port

	fmt.Printf("server running on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
