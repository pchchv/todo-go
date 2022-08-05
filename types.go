package main

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Completed bool   `json:"completed"`
}
