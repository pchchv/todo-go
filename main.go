package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"textx"`
	Completed bool   `json:"completed"`
}

var (
	envURL  string
	testURL string
)

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing.
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func creator(title string, text string, completed string) Todo {
	todo := Todo{}
	todo.Title = title
	todo.Text = text
	// TODO: Implement id generation and assignment
	if completed == "true" {
		todo.Completed = true
	} else {
		todo.Completed = false
	}
	// TODO: Implement addition of a task repository
	return todo
}

func getter(id string, title string) (Todo, []Todo) {
	var todos []Todo
	todo := Todo{}
	if id != "" {
		// TODO: Implement a getting task by id
	} else if title != "" {
		// TODO: Implement a getting task by title
	} else {
		// TODO: Implement a getting all tasks
	}
	return todo, todos
}

func patcher(id string, title string, text string, completed string) Todo {
	t, _ := getter(id, title)
	// TODO: Implement a task updating
	return t
}

func deleter(id string, title string) (Todo, []Todo) {
	var todos []Todo
	var todo Todo
	if id != "" {
		// TODO: Implement task deleting by id
	} else if title != "" {
		// TODO: Implement task deleting by title
	} else {
		// TODO: Implement deleting all tasks
	}
	return todo, todos
}

func main() {
	envURL = getEnvValue("HOST") + ":" + getEnvValue("PORT")
	server()
}
