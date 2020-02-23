package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"golang.org/x/sync/errgroup"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	box, err := rice.FindBox("../client/dist")
	if err != nil {
		return fmt.Errorf("error opening rice.Box: %w", err)
	}

	http.Handle("/", http.FileServer(box.HTTPBox()))

	http.HandleFunc("/hello", sendMessage())

	eg := errgroup.Group{}
	eg.Go(func() error {
		return http.ListenAndServe(":8000", nil)
	})

	return eg.Wait()
}

func sendMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("incoming request from %s", r.UserAgent())
		message := Message{"John Smith"}

		data, err := json.Marshal(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		_, _ = w.Write(data)
	}

}
