package cmd

import (
	"fmt"
	"todo-cli/data"
	"todo-cli/utils"
)

func Add(title string, todos *[]data.Task) {
	task := data.CreateTask(title, len(*todos))

	*todos = append(*todos, *task)

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
		*allTodo = utils.RemoveArrayElementOnIndex(indexToBeDeleted, *allTodo)

		fmt.Println("Task is removed, below is the remaining task")
		fmt.Println("============================================")

		for _, element := range *allTodo {
			fmt.Printf("%d - %s\n", element.Id, element.Title)
		}

	}
}
