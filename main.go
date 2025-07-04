package main

import (
	"fmt"

	"github.com/gabsgasps/todo-list-golang/cmd"
)

func main() {
	todos := cmd.Todos{}

	storage := cmd.NewStorage[cmd.Todos]("todos.json")
	err := storage.Load(&todos)
	if err != nil {
		fmt.Println("Warning: Could not load todos from storage. Starting fresh todos.")
		todos = cmd.Todos{}
	}

	cmdFlags := cmd.ParseFlags()
	cmdFlags.Execute(&todos)

	// Save to storage
	err = storage.Save(todos)
	if err != nil {
		fmt.Printf("Error saving todos in storage: %v\n", err)
	}
}
