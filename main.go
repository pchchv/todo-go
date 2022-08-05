package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envURL string
var testURL string

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

func main() {
	envURL = getEnvValue("HOST") + ":" + getEnvValue("PORT")
	server()
}
