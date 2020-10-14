package main

import (
  "log"
  "net/http"
  "os"

	"github.com/aws/aws-xray-sdk-go/xray"
)

func main() {
  http.Handle("/hi", xray.Handler(xray.NewFixedSegmentNamer("xray-k8s-app"), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
  })))
  http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("works!"))
  })
  http.HandleFunc("/die", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("exiting.."))
    os.Exit(1)
  })
  log.Println("about to listen on port 8989")
  log.Fatal(http.ListenAndServe(":8989", nil))
}
