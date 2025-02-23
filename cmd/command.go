package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"todo-cli/data"
)

func Add(title string, todos *[]data.Task) {
	task := data.CreateTask(title, len(*todos))

	*todos = append(*todos, *task)

	rawJson, err := json.Marshal(todos)
	if err != nil {
		log.Fatal("Can't create new task")
	}

	err = os.WriteFile("todo.json", rawJson, 0600)
	if err != nil {
		log.Fatal("Failed to write the data in file")
	}

	fmt.Println("<<Task has been created>>")

}

func ListTasks(todos *[]data.Task) {
	fmt.Println("Here is the list of all the task")
	fmt.Println("________________________________")
	for _, element := range *todos {
		fmt.Printf("%d - %s\n", element.Id, element.Title)
	}
}

func DeleteTask(id uint32, allTodo *[]data.Task) {
	var indexToBeDeleted int

	for index, val := range *allTodo {
		if val.Id == uint32(id) {
			indexToBeDeleted = index
		}
	}

	if indexToBeDeleted >= 0 {
		*allTodo = append((*allTodo)[:indexToBeDeleted], (*allTodo)[indexToBeDeleted+1:]...)

		fmt.Println("Task is removed, below is the remaining task")
		fmt.Println("============================================")

		for _, element := range *allTodo {
			fmt.Printf("%d - %s\n", element.Id, element.Title)
		}

	}
}
