package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var flag string
	var task string

	args := os.Args[1:]
	if len(args) >= 2 {
		task = args[1]

	}
	flag = args[0]

	var allTodo []Task
	data, err := os.ReadFile("todo.json")

	if err != nil {
		log.Fatal("Failed to load file")
	}

	err = json.Unmarshal(data, &allTodo)

	if err != nil {
		log.Fatal("Failed to load file")
	}

	switch flag {

	case "list":
		fmt.Println("Here is the list of all the task")
		fmt.Println("________________________________")
		for _, element := range allTodo {
			fmt.Printf("%d - %s\n", element.Id, element.Title)
		}

	case "add":

		task := createNewTask(task, len(allTodo))
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

	case "delete":
		var indexToBeDeleted int

		for index, val := range allTodo {
			id, _ := strconv.ParseUint(task, 10, 32)

			if val.Id == uint32(id) {
				indexToBeDeleted = index
			}
		}

		if indexToBeDeleted >= 0 {
			allTodo = append(allTodo[:indexToBeDeleted], allTodo[indexToBeDeleted+1:]...)

			fmt.Println("Task is removed, below is the remaining task")
			fmt.Println("============================================")

			for _, element := range allTodo {
				fmt.Printf("%d - %s\n", element.Id, element.Title)
			}

		}

	}

}
