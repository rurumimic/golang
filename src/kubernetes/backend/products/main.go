package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	products := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince", "raspberry", "strawberry", "tangerine", "ugli", "watermelon"}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().UnixNano() % 200
		time.Sleep(time.Duration(t) * time.Millisecond)

		w.Write([]byte(products[t%int64(len(products))]))
		w.Write([]byte("\n"))
	})
	http.ListenAndServe(":" + port, r)
}
