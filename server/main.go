package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/indiependente/pkg/logger"
	"golang.org/x/sync/errgroup"
)

const serviceName = `backend`

type Message struct {
	Text string `json:"text"`
}

func main() {
	//ctx := context.Background()
	logger := logger.GetLoggerString(serviceName, os.Getenv("LOG_LEVEL"))
	if err := run(logger); err != nil {
		logger.Fatal("Error while running", err)
	}
}

func run(logger logger.Logger) error {
	box, err := rice.FindBox("../client/dist")
	if err != nil {
		return fmt.Errorf("error opening rice.Box: %w", err)
	}

	http.Handle("/", http.FileServer(box.HTTPBox()))

	http.HandleFunc("/hello", requestIDMiddleware(sendMessage(logger)))

	eg := errgroup.Group{}
	eg.Go(func() error {
		return http.ListenAndServe(":8000", nil)
	})

	return eg.Wait()
}

func sendMessage(logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := r.Context().Value("reqID").(string)
		logger.RequestID(strfmt.UUID(id)).Info("incoming request from " + r.UserAgent())

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

func requestIDMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		ctx := context.WithValue(context.Background(), "reqID", uuid.String())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
