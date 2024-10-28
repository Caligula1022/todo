package main

type Todo struct {
	ID     int
	Title  string
	Status string
}

type TodoList struct {
	Todos []Todo
}
