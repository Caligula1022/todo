package main

import (
	"fmt"
	"os"
)

const filename = "todo.json"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	command := os.Args[1]

	// loda todo files
	todolist := &TodoList{}
	err := todolist.Load(filename)
	if err != nil {
		fmt.Printf("loading todo file failed: %v", err)
		return
	}

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("there should be a title")
			return
		}
		title := os.Args[2]
		todolist.Add(title)
	case "list":
	case "done":
	case "delete":
	default:
		printUsage()
		return
	}

	err = todolist.Save(filename)
	if err != nil {
		fmt.Printf("save file failed: %v", err)
		return
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("	todo add [title]")
	fmt.Println("	todo list")
	fmt.Println("	todo done [id]")
	fmt.Println("	todo delete [id]")
}
