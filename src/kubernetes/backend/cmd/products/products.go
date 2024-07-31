package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	products := []string{"APPLE", "BANANA", "CHERRY", "DATE", "ELDERBERRY", "FIG", "GRAPE", "HONEYDEW", "KIWI", "LEMON", "MANGO", "NECTARINE", "ORANGE", "PAPAYA", "QUINCE", "RASPBERRY", "STRAWBERRY", "TANGERINE", "UGLI", "WATERMELON"}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().UnixNano() % 100
		time.Sleep(time.Duration(t) * time.Millisecond)

		answer := []byte(products[t%int64(len(products))])
		fmt.Println("Answer: ", string(answer))
		w.Write(answer)
	})
	http.ListenAndServe(":"+port, r)
}
