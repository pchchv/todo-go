package main

import (
	"errors"
	"log"
	"os"
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
	// Creates a new todo and adds it to the repository
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
	// Gets the todo by id or title, or gets a list of all todos
	var todos []*Todo
	todo := new(Todo)
	if id != "" {
		nid, err := strconv.Atoi(id)
		if err != nil {
			log.Panic(err)
		}
		todo, err := todoRepository.Get(nid)
		if err != nil {
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
	// Changes the todo by finding it by id or title
	t, _, _ := getter(id, title)
	// TODO: Implement a task updating
	return t
}

func deleter(id string, title string) error {
	// Deletes the todo by id or title, or deletes all todos
	if id != "" {
		nid, err := strconv.Atoi(id)
		if err != nil {
			log.Panic(err)
		}
		err = todoRepository.Delete(nid)
		if err != nil {
			return err
		}
	} else if title != "" {
		// TODO: Implement task deleting by title
	} else {
		todoRepository.DeleteAll()
	}
	return nil
}

func main() {
	envURL = getEnvValue("HOST") + ":" + getEnvValue("PORT")
	server()
}
