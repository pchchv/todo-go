package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func pingHandler(c echo.Context) error {
	msg := "Todo backend Service. Version 0.2.0"
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
	todo, _, err := getter(id, title)
	if err != nil {
		return c.String(http.StatusNotFound, "Todo not found")
	}
	return c.JSONPretty(http.StatusOK, todo, "\t")
}

func patchTodoHandler(c echo.Context) error {
	id := c.QueryParam("id")
	title := c.QueryParam("title")
	text := c.QueryParam("text")
	completed := c.QueryParam("completed")
	todo, err := updater(id, title, text, completed)
	if err != nil {
		return c.String(http.StatusNotFound, "Todo not found")
	}
	return c.JSONPretty(http.StatusOK, todo, "\t")
}

func deleteTodoHandler(c echo.Context) error {
	id := c.QueryParam("id")
	title := c.QueryParam("title")
	err := deleter(id, title)
	if err != nil {
		return c.String(http.StatusNotFound, "Todo not found")
	}
	return c.NoContent(http.StatusNoContent)
}

func deleteAllTodosHandler(c echo.Context) error {
	deleter("", "")
	return c.NoContent(http.StatusNoContent)
}

func getAllTodosHandler(c echo.Context) error {
	_, todos, _ := getter("", "")
	return c.JSONPretty(http.StatusOK, todos, "\t")
}

func routes(e *echo.Echo) {
	e.GET("/", pingHandler)
	e.GET("/ping", pingHandler)
	e.POST("/todo", createTodoHandler)
	e.GET("/todo", getTodoHandler)
	e.PATCH("/todo", patchTodoHandler)
	e.DELETE("/todo", deleteTodoHandler)
	e.GET("/todos", getAllTodosHandler)
	e.DELETE("/todos", deleteAllTodosHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(":" + getEnvValue("PORT")))
}
