package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

type Task struct {
	Id          uint32 `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

func createNewTask(title string) *Task {
	return &Task{Id: uuid.New().ID(), Title: title, IsCompleted: false}

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

	var i int
	fmt.Println("What do you want to do?")
	fmt.Println("1. List all task")
	fmt.Println("2. Create new task")
	fmt.Println("3. Delete task of id: ")
	fmt.Scanln(&i)

	switch i {
	case 1:
		fmt.Println("Here is the list of all the task")
		for _, element := range allTodo {
			fmt.Printf("%d - %s\n", element.Id, element.Title)
		}

	case 2:
		fmt.Println("create new task")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		title := scanner.Text()

		task := createNewTask(title)
		allTodo = append(allTodo, *task)

		rawJson, err := json.Marshal(allTodo)
		if err != nil {
			log.Fatal("Can't create new task")
		}

		err = os.WriteFile("todo.json", rawJson, 0600)
		if err != nil {
			log.Fatal("Failed to write the data in file")
		}

	case 3:
		fmt.Println("delete new task")

	default:
		fmt.Println("Invalid response")

	}

}
