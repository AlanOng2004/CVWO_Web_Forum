package main

import (
    "fmt"
    "net/http"

    "forum-backend/internal/repositories"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
)

func main() {
    db := repositories.NewDB("forum.db")
    repositories.RunMigrations(db, "migrations")

    r := chi.NewRouter()

    // Basic middleware
    r.Use(middleware.Logger)

    r.Use(cors.Handler(cors.Options{
        // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
        AllowedOrigins:   []string{"http://localhost:5173"},
        // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
      }))

    // Health endpoint
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status":"ok"}`))
    })

    port := 8080
    fmt.Printf("Server running on port %d\n", port)
    http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

