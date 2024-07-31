package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Result struct {
	Response *http.Response
	Error    error
	Url      string
}

type Content struct {
	Url  string
	Body string
	Ok   bool
}

func AsyncGet(ctx context.Context, urls ...string) <-chan Result {
	results := make(chan Result, len(urls))

	go func() {
		defer close(results)

		for _, url := range urls {
			req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
			client := &http.Client{}
			res, err := client.Do(req)

			select {
			case <-ctx.Done():
        if res != nil {
          res.Body.Close()
        }
				return
			case results <- Result{res, err, url}:
			}
		}

	}()

	return results
}

func main() {
	urls := []string{"http://localhost:3001", "http://localhost:3002"}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Millisecond)
		defer cancel()

		contents := []Content{}
		for result := range AsyncGet(ctx, urls...) {
			if result.Error != nil {
				contents = append(contents, Content{Url: result.Url, Body: result.Error.Error(), Ok: false})
				continue
			}
			body, err := io.ReadAll(result.Response.Body)
			result.Response.Body.Close()
			if err != nil {
				fmt.Println(err)
				contents = append(contents, Content{Url: result.Url, Body: err.Error(), Ok: false})
				continue
			}
			contents = append(contents, Content{Url: result.Url, Body: string(body), Ok: true})
		}

		for _, content := range contents {
			if content.Ok {
				w.Write([]byte("OK: "))
			} else {
				w.Write([]byte("NG: "))
			}

			w.Write([]byte(content.Url + " : " + content.Body + "\n"))
		}
	})

	http.ListenAndServe(":3000", r)
}
