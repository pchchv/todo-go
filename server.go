package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func pingHandler(c echo.Context) error {
	msg := "Todo backend Service. Version 0.0.0"
	return c.String(http.StatusOK, msg)
}

func createTodoHandler(c echo.Context) error {
	todo := Todo{}
	todo.Title = c.QueryParam("title")
	todo.Text = c.QueryParam("text")
	if c.QueryParam("completed") == "true" {
		todo.Completed = true
	} else {
		todo.Completed = false
	}
	return c.JSONPretty(http.StatusOK, creator(todo), "\t")
}

func routes(e *echo.Echo) {
	e.GET("/ping", pingHandler)
	e.GET("/todo", createTodoHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(envURL))
}
