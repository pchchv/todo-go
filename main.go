package main

import (
	"log"
	"os"
	"errors"
	"strconv"

	"github.com/joho/godotenv"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"textx"`
	Completed bool   `json:"completed"`
}

var (
	envURL         string
	testURL        string
	todoRepository = NewInMemoryTodoRepository()
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

func creator(title string, text string, completed string) *Todo {
	// Creates a new task and adds it to the repository
	todo := new(Todo)
	todo.Title = title
	todo.Text = text
	if completed == "true" {
		todo.Completed = true
	} else {
		todo.Completed = false
	}
	todoRepository.Create(todo)
	return todo
}

func getter(id string, title string) (*Todo, []*Todo, error) {
	var todos []*Todo
	todo := new(Todo)
	if id != "" {
		nid, err := strconv.Atoi(id)
		if err != nil {
			log.Panic(err)
		}
		if todo, err := todoRepository.Get(nid); err != nil {
			return todo, todos, errors.New("Todo not found")
		}
	} else if title != "" {
		// TODO: Implement a getting task by title
	} else {
		todos = todoRepository.GetAll()
	}
	return todo, todos, nil
}

func patcher(id string, title string, text string, completed string) *Todo {
	t, _, _ := getter(id, title)
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
