package main

import (
  "context"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-xray-sdk-go/xray"
)

func exit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("exiting.."))
		os.Exit(1)
	}
}

func getIP(c *http.Client, url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    req, _ := http.NewRequest("GET", url, nil)
    ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second) // http ttl
    defer cancel()
		res, err := c.Do(req.WithContext(ctx))
		if err != nil {
      log.Printf("ERROR: request to %s failed: %s\n", url, err)
			http.Error(w, "server error", 500)
			return
		}
		defer res.Body.Close()
		io.Copy(w, res.Body) // TODO: forward http status
	}
}

func health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}
}

func main() {
	url := os.Getenv("AWS_LAMBDA_URL")
	if url == "" {
		log.Println("env var AWS_LAMBDA_URL missing")
		os.Exit(1)
	}
	client := xray.Client(&http.Client{
		Timeout: 5 * time.Second, // tcp ttl
	})
	http.Handle("/die", exit())
	http.Handle("/ip", xray.Handler(xray.NewFixedSegmentNamer("xray-k8s-app"), getIP(client, url)))
	http.Handle("/health", health())
	log.Println("about to listen on port 8989")
	log.Fatal(http.ListenAndServe(":8989", nil))
}
