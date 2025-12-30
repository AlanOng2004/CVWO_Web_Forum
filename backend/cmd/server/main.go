package main

import (
    "fmt"
    "net/http"

    "forum-backend/internal/repositories"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    db := repositories.NewDB("forum.db")
    repositories.RunMigrations(db, "migrations")

    r := chi.NewRouter()

    // Basic middleware
    r.Use(middleware.Logger)

    // Health endpoint
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status":"ok"}`))
    })

    port := 8080
    fmt.Printf("Server running on port %d\n", port)
    http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

