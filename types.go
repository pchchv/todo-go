package main

type Todo struct {
	Title     string `json:"title"`
	Text      string `json:"textx"`
	Completed bool   `json:"completed"`
}

type MongoTodo struct {
	Id   string `json:"id"`
	Todo Todo
}
