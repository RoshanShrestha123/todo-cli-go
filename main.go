package main

import (
	"os"
	"strconv"

	"todo-cli/cmd"
	"todo-cli/data"
)

func main() {

	var flag string
	var todo string

	args := os.Args[1:]
	if len(args) >= 2 {
		todo = args[1]

	}
	flag = args[0]
	var allTodo []data.Task
	data.ReadDataFromFile(&allTodo)

	switch flag {
	case "list":
		cmd.ListTasks(&allTodo)

	case "add":
		cmd.Add(todo, &allTodo)

	case "delete":
		id, _ := strconv.ParseUint(todo, 10, 32)
		cmd.DeleteTask(uint32(id), &allTodo)

	}

	data.SaveDataToFile(&allTodo)

}
