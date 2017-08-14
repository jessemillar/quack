package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jessemillar/health"
	"github.com/jessemillar/quack/quack"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	port := flag.String("p", "8000", "port to serve on")
	flag.Parse()

	router.GET("/", echo.WrapHandler(http.HandlerFunc(quack.Quack)))
	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	router.GET("/quack", echo.WrapHandler(http.HandlerFunc(quack.Quack)))

	log.Printf("Listening on HTTP port: %s\n", *port)

	server := http.Server{
		Addr:           ":" + *port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
