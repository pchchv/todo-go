package main

import (
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
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}
func creator(todo Todo) MongoTodo {
	t := MongoTodo{}
	t.Todo = todo
	_, err := bson.Marshal(t)
	if err != nil {
		log.Panic(err)
	}
	// TODO: Implement addition of a task to the database
	return t
}

func main() {
	envURL = getEnvValue("HOST") + ":" + getEnvValue("PORT")
	db()
	server()
}
