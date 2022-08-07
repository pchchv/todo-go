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
	title := c.QueryParam("title")
	text := c.QueryParam("text")
	completed := c.QueryParam("completed")
	todo := creator(title, text, completed)
	return c.JSONPretty(http.StatusOK, todo, "\t")
}

func getTodoHandler(c echo.Context) error {
	id := c.QueryParam("id")
	title := c.QueryParam("title")
	todo := getter(id, title)
	return c.JSONPretty(http.StatusOK, todo, "\t")
}

func patchTodoHandler(c echo.Context) error {
	id := c.QueryParam("id")
	title := c.QueryParam("title")
	text := c.QueryParam("text")
	completed := c.QueryParam("completed")
	todo := patcher(id, title, text, completed)
	return c.JSONPretty(http.StatusOK, todo, "\t")
}

func deleteTodoHandler(c echo.Context) error {
	id := c.QueryParam("id")
	title := c.QueryParam("title")
	todo, _ := deleter(id, title)
	return c.JSONPretty(http.StatusOK, todo, "\t")
}

func deleteAllTodosHandler(c echo.Context) error {
	_, todos := deleter("", "")
	return c.JSONPretty(http.StatusOK, todos, "\t")
}

func routes(e *echo.Echo) {
	e.GET("/", pingHandler)
	e.GET("/ping", pingHandler)
	e.POST("/todo", createTodoHandler)
	e.GET("/todo", getTodoHandler)
	e.PATCH("/todo", patchTodoHandler)
	e.DELETE("/todo", deleteTodoHandler)
	e.DELETE("/todos", deleteAllTodosHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(envURL))
}
