package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Task struct {
	Id          uint32 `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

func createNewTask(title string, len int) *Task {
	return &Task{Id: uint32(len), Title: title, IsCompleted: false}

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

	for {
		var i int
		fmt.Println("What do you want to do?")
		fmt.Println("=========================")
		fmt.Println("1. List all task")
		fmt.Println("2. Create new task")
		fmt.Println("3. Delete task of id: ")
		fmt.Scanln(&i)

		switch i {
		case 1:
			fmt.Println("Here is the list of all the task")
			fmt.Println("________________________________")
			for _, element := range allTodo {
				fmt.Printf("%d - %s\n", element.Id, element.Title)
			}

		case 2:
			fmt.Println("create new task")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			title := scanner.Text()

			task := createNewTask(title, len(allTodo))
			allTodo = append(allTodo, *task)

			rawJson, err := json.Marshal(allTodo)
			if err != nil {
				log.Fatal("Can't create new task")
			}

			err = os.WriteFile("todo.json", rawJson, 0600)
			if err != nil {
				log.Fatal("Failed to write the data in file")
			}

			fmt.Println("<<Task has been created>>")

		case 3:
			fmt.Println("delete new task")

		default:
			fmt.Println("Invalid response")

		}

		fmt.Println("You want to continue? Y/n")
		var c string
		fmt.Scan(&c)

		if c == "y" || c == "Y" {
			continue
		} else {
			return
		}
	}

}
