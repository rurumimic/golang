package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
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

func FanIn(ctx context.Context, channels ...<-chan Result) <-chan Result {
	var wg sync.WaitGroup
	multiplexed := make(chan Result)

	multiplex := func(c <-chan Result) {
		defer wg.Done()

		for i := range c {
			select {
			case <-ctx.Done():
				log.Printf("FanIn: %s %s\n", i.Url, ctx.Err().Error())
				multiplexed <- i
			case multiplexed <- i:
				log.Printf("Multiplexed: %s\n", i.Url)
			}
		}
	}

	wg.Add(len(channels))

	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexed)
	}()

	return multiplexed
}

func FetchData(ctx context.Context, url string) <-chan Result {
	stream := make(chan Result)

	go func() {
		defer close(stream)

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			log.Printf("Error new request: %s\n", err.Error())
			stream <- Result{nil, err, url}
			return
		}

		client := &http.Client{}
		res, err := client.Do(req)

		select {
		case <-ctx.Done():
			if res != nil {
				res.Body.Close()
			}
			if ctx.Err() == context.DeadlineExceeded {
				log.Printf("WithTimeout: %s %s\n", url, ctx.Err().Error())
				stream <- Result{nil, ctx.Err(), url}
			} else {
				log.Printf("Error fetching data: %s %s\n", url, ctx.Err().Error())
				stream <- Result{nil, ctx.Err(), url}
			}
			return
		default:
			if err != nil {
				log.Printf("Error fetching data! %s %s\n", url, err)
				stream <- Result{nil, err, url}
				return
			}

			log.Printf("Response status: %s\n", res.Status)
			stream <- Result{res, err, url}
		}
		log.Printf("Stream done\n")
	}()

	return stream
}

func AsyncGet(ctx context.Context, urls ...string) []Content {
	results := make([]<-chan Result, len(urls))
	contents := []Content{}

	for i, url := range urls {
		results[i] = FetchData(ctx, url)
	}

	for result := range FanIn(ctx, results...) {
		if result.Error != nil {
			contents = append(contents, Content{Url: result.Url, Body: result.Error.Error(), Ok: false})
			continue
		}

		body, err := io.ReadAll(result.Response.Body)
		result.Response.Body.Close()

		if err != nil {
			log.Printf("Error reading data: %s\n", err.Error())
			contents = append(contents, Content{Url: result.Url, Body: err.Error(), Ok: false})
			continue
		}

		contents = append(contents, Content{Url: result.Url, Body: string(body), Ok: true})
	}

	log.Printf("Contents: %v\n", contents)

	return contents
}

func main() {
	urls := []string{"http://localhost:3001", "http://localhost:3002"}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		contents := AsyncGet(ctx, urls...)

		response := map[string][]Content{
			"results": contents,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
