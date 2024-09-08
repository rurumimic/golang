package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const defaultPort = "3001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	users := []string{"john", "jane", "doe", "alice", "bob", "charlie", "david", "eve", "frank", "grace", "heidi"}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().UnixNano() % 1000
		time.Sleep(time.Duration(t) * time.Millisecond)

		answer := []byte(users[t%int64(len(users))])
		fmt.Println("Answer: ", string(answer))
		w.Write(answer)
	})
	http.ListenAndServe(":"+port, r)
}
