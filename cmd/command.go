package cmd

import (
	"errors"
	"fmt"
	"todo-cli/data"
	"todo-cli/utils"
)

func Add(title string, todos *[]data.Task) {

	maxId := 0

	for _, v := range *todos {
		if v.Id > uint32(maxId) {
			maxId = int(v.Id)
		}
	}
	task := data.CreateTask(title, uint32(maxId+1))

	*todos = append(*todos, *task)

	fmt.Println("<<Task has been created>>")

}

func ListTasks(todos *[]data.Task, status string) {
	fmt.Println("--Tasks--")
	fmt.Println("________________________________")
	for _, element := range *todos {
		if element.Status == status || status == "" {

			fmt.Printf("%d - %s - %s\n", element.Id, element.Title, element.Status)
		}
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

func UpdateTask(id uint32, value data.Task, allTodo *[]data.Task) {

	var indexToBeDeleted int

	for index, val := range *allTodo {
		if val.Id == uint32(id) {
			indexToBeDeleted = index
		}
	}
	(*allTodo)[indexToBeDeleted] = value
}

func SelectTaskFromId(id uint32, allTodo *[]data.Task) (*data.Task, error) {

	for _, val := range *allTodo {
		if val.Id == uint32(id) {
			return &val, nil
		}
	}
	return nil, errors.New("No data found")
}
