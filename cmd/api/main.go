package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/OnatArslan/go-rest-blog/internal/config"
	"github.com/OnatArslan/go-rest-blog/internal/httpx"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Load config files
	cfg := config.Load()

	// Create context in here with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	// This block for connect postgres with native pgx
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	log.Println("Database connected...")

	// this db will used in sqlc repositories
	// db := db.New(pool)

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

// Example sqlc and pgx setup
/*
func main() {
    // urlExample := "postgres://username:password@localhost:5432/database_name"
    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
            fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
            os.Exit(1)
    }
    defer conn.Close(context.Background())

    q := db.New(conn)

    author, err := q.GetAuthor(context.Background(), 1)
    if err != nil {
            fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
            os.Exit(1)
    }

    fmt.Println(author.Name)
*/
