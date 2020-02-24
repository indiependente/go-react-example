package main

import (
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	//ctx := context.Background()
	//logger := logger.GetLoggerString(serviceName, os.Getenv("LOG_LEVEL"))
	if err := run(); err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run() error {
	e := echo.New()
	//e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	assetHandler := http.FileServer(rice.MustFindBox("../client/dist").HTTPBox())
	// the file server serves the index.html from the rice box
	e.GET("/", echo.WrapHandler(assetHandler))
	// servers other static files
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", assetHandler)))
	e.GET("/hello", sendMessage())
	return e.Start(":8000")
	//return e.StartAutoTLS(":443")
}

func sendMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		message := Message{"John Smith"}
		return c.JSON(200, message)
	}
}
