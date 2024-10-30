package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TodoList struct {
	Todos []Todo `json:"todos"`
}

func (t *TodoList) Load(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Todos = []Todo{}
		return nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file failed: %v", err)
		return err
	}
	// if len(data) == 0 {
	// 	t.Todos = []Todo{}
	// 	return nil
	// 	// err = t.Save(filename)
	// 	// if err != nil {
	// 	// 	fmt.Printf("save file failed: %v", err)
	// 	// 	return err
	// 	// }
	// }
	err = json.Unmarshal(data, t)
	if err != nil {
		fmt.Println("JSON 解析失败，初始化为空列表。")
		t.Todos = []Todo{}
		return t.Save(filename)
	}
	return nil
}

// save todo file
func (t *TodoList) Save(filename string) error {
	data, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		fmt.Printf("Save todo file failed: %v", err)
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// add todo item
func (t *TodoList) Add(title string) error {
	id := 1
	if len(t.Todos) > 0 {
		id = len(t.Todos) + 1
	}
	todo := Todo{ID: id, Title: title, Status: "pending"}
	t.Todos = append(t.Todos, todo)
	return nil
}

// list all todo item

// make todo item done

// delete todo item
