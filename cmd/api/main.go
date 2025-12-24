package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/OnatArslan/go-rest-blog/internal/config"
	"github.com/OnatArslan/go-rest-blog/internal/db"
	"github.com/OnatArslan/go-rest-blog/internal/httpx"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Load config files
	cfg := config.Load()

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	log.Println("Database connected...")

	// this db will used in sqlc repositories
	db := db.New(pool)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	log.Println("listening on", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, r))
}
