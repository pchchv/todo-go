package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	envURL     string
	testURL    string
	collection *mongo.Collection
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

func creator(title string, text string, completed string) MongoTodo {
	todo := Todo{}
	t := MongoTodo{}
	todo.Title = title
	todo.Text = text
	if completed == "true" {
		todo.Completed = true
	} else {
		todo.Completed = false
	}
	t.Todo = todo
	_, err := bson.Marshal(t)
	if err != nil {
		log.Panic(err)
	}
	// TODO: Implement addition of a task to the database
	return t
}

func getter(id string, title string) (MongoTodo, error) {
	var res *mongo.SingleResult
	result := &MongoTodo{}
	if id != "" {
		res = collection.FindOne(context.TODO(), bson.M{"_id": id})
	} else {
		res = collection.FindOne(context.TODO(), bson.M{"title": title})
	}
	err := res.Decode(result)
	if err != nil {
		return *result, errors.New("Todo not found")
	}
	return *result, nil
}

func patcher(id string, title string, text string, completed string) MongoTodo {
	t, err := getter(id, title)
	if err != nil {
		log.Panic(err)
	}
	// TODO: Implement a task updating
	return t
}

func main() {
	envURL = getEnvValue("HOST") + ":" + getEnvValue("PORT")
	db()
	server()
}
