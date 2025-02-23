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

func SaveDataToFile(todos *[]Task) {
	rawJson, err := json.Marshal(todos)
	if err != nil {
		log.Fatal("Can't create new task")
	}

	err = os.WriteFile("todo.json", rawJson, 0600)
	if err != nil {
		log.Fatal("Failed to write the data in file")
	}
}
