package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ping(c echo.Context) error {
	msg := "Todo backend Service. Version 0.0.0"
	return c.String(http.StatusOK, msg)
}

func routes(e *echo.Echo) {
	e.GET("/ping", ping)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(envURL))
}
