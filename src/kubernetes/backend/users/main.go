package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	users := []string{"John", "Jane", "Doe", "Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Heidi"}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().UnixNano() % 100
		time.Sleep(time.Duration(t) * time.Millisecond)

		w.Write([]byte(users[t%int64(len(users))]))
	})
	http.ListenAndServe(":"+port, r)
}
