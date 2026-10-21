package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"moviehub/handler"
	"moviehub/mw"
	"moviehub/store"
)

func main() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("ต่อฐานข้อมูลไม่ติด ", err)
	}

	h := &handler.Handler{Store: &store.Store{Pool: pool}}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/auth/register", h.Register)
	mux.HandleFunc("GET /api/movies", h.SearchMovies)
	mux.HandleFunc("GET /api/movies/{id}/reviews", h.ListReviews)

	log.Println("MovieHub API ฟังอยู่ที่ :8080")
	log.Fatal(http.ListenAndServe(":8080", mw.Logging(mux)))
}