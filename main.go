package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

func main() {

	var allTodo []Task
	data, err := os.ReadFile("todo.json")

	if err != nil {
		log.Fatal("Failed to load file")
	}

	err = json.Unmarshal(data, &allTodo)

	if err != nil {
		log.Fatal("Failed to load file")
	}

	for _, element := range allTodo {
		fmt.Printf("%d - %s\n", element.Id, element.Title)
	}

}
