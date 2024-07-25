package main

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		urls := []string{
			"http://localhost:3001",
			"http://localhost:3002",
		}

		var wg sync.WaitGroup
		var mu sync.Mutex
		responses := make([]*http.Response, 0)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		for _, url := range urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
				client := &http.Client{}
				resp, err := client.Do(req)
				if err == nil {
					mu.Lock()
					responses = append(responses, resp)
					mu.Unlock()
				}
			}(url)
		}

		wg.Wait()

		for _, resp := range responses {
			defer resp.Body.Close()
			w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
			w.WriteHeader(resp.StatusCode)
			io.Copy(w, resp.Body)
		}

		if len(responses) == 0 {
			http.Error(w, "No responses received within 100ms", http.StatusGatewayTimeout)
		}
	})
	http.ListenAndServe(":3000", r)
}
