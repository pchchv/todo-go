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
	Text      string `json:"text"`
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
	var err error
	var todos []*Todo
	todo := new(Todo)
	if id != "" {
		nid, err := strconv.Atoi(id)
		if err != nil {
			log.Panic(err)
		}
		todo, err = todoRepository.GetById(nid)
		if err != nil {
			return todo, todos, errors.New("Todo not found")
		}
	} else if title != "" {
		todo, err = todoRepository.GetByTitle(title)
		if err != nil {
			return todo, todos, errors.New("Todo not found")
		}
	} else {
		todos = todoRepository.GetAll()
	}
	return todo, todos, nil
}

func updater(id string, title string, text string, completed string) (*Todo, error) {
	// Changes the todo by finding it by id or title
	todo, _, err := getter(id, title)
	if err != nil {
		return todo, err
	}
	todo.Title = title
	todo.Text = text
	if completed == "true" {
		todo.Completed = true
	} else if completed == "false" {
		todo.Completed = false
	}
	err = todoRepository.Update(todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
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
		todo, _, err := getter("", title)
		if err != nil {
			return err
		}
		err = todoRepository.Delete(todo.Id)
		if err != nil {
			return err
		}
	} else {
		todoRepository.DeleteAll()
	}
	return nil
}

func main() {
	envURL = getEnvValue("HOST") + ":" + getEnvValue("PORT")
	server()
}
