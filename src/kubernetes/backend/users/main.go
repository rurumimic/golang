package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
  "os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    t := time.Now().UnixNano() % 200
    time.Sleep(time.Duration(t) * time.Millisecond)

		w.Write([]byte(`[{"name":"John Doe","age":25},{"name":"Jane Doe","age":24}]`))
	})
	http.ListenAndServe(":3001", r)
}
