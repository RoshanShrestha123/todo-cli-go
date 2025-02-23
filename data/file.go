package data

import (
	"encoding/json"
	"log"
	"os"
)

func ReadDataFromFile(todos *[]Task) {
	data, err := os.ReadFile("todo.json")

	if err != nil {
		log.Fatal("Failed to load file")
	}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		log.Fatal("Failed to load file")
	}
}
