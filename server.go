package main

import (
	"net/http"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/jessemillar/health"
	"github.com/jessemillar/quack/quack"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	port := kingpin.Flag("port", "The port to listen on.").Short('p').Default("8000").String()
	kingpin.Parse()

	router.GET("/", echo.WrapHandler(http.HandlerFunc(quack.Quack)))
	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	router.GET("/quack", echo.WrapHandler(http.HandlerFunc(quack.Quack)))
	router.POST("/quack", echo.WrapHandler(http.HandlerFunc(quack.Quack)))
	router.PUT("/quack", echo.WrapHandler(http.HandlerFunc(quack.Quack)))
	router.DELETE("/quack", echo.WrapHandler(http.HandlerFunc(quack.Quack)))

	server := http.Server{
		Addr:           ":" + *port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
