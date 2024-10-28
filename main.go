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

	switch command {
	case "add":
	case "list":
	case "done":
	case "delete":
	default:
		printUsage()
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
