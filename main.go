package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"todo-cli/cmd"
	"todo-cli/data"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: <command> <todo> [title]")
		os.Exit(1)
	}

	flag := args[0]
	todo, title := "", ""

	if len(args) > 1 {
		todo = args[1]
	}
	if len(args) > 2 {
		title = args[2]
	}
	if len(args) > 3 {
		fmt.Println("Error: Too many arguments. Expected format: <command> <todo> [title]")
		os.Exit(1)
	}

	var allTodo []data.Task
	data.ReadDataFromFile(&allTodo)

	switch flag {
	case "list":

		cmd.ListTasks(&allTodo, todo)

	case "add":
		cmd.Add(todo, &allTodo)

		fmt.Printf("-- %s is added successfully --\n", todo)

	case "delete":
		id, _ := strconv.ParseUint(todo, 10, 32)
		cmd.DeleteTask(uint32(id), &allTodo)

		fmt.Printf("-- %d is deleted successfully --\n", id)

	case "update":
		id, _ := strconv.ParseUint(todo, 10, 32)
		task, err := cmd.SelectTaskFromId(uint32(id), &allTodo)
		if err != nil {
			log.Fatal(err)
		}

		task.Title = title
		cmd.UpdateTask(uint32(id), *task, &allTodo)

		fmt.Printf("-- %d is updated successfully --\n", task.Id)

	case "mark-done":
		id, _ := strconv.ParseUint(todo, 10, 32)
		task, err := cmd.SelectTaskFromId(uint32(id), &allTodo)
		if err != nil {
			log.Fatal(err)
		}

		task.Status = "done"
		cmd.UpdateTask(uint32(id), *task, &allTodo)
		fmt.Printf("-- %s is marked as done --\n", task.Title)

	case "mark-in-progress":
		id, _ := strconv.ParseUint(todo, 10, 32)
		task, err := cmd.SelectTaskFromId(uint32(id), &allTodo)
		if err != nil {
			log.Fatal(err)
		}

		task.Status = "in-progress"
		cmd.UpdateTask(uint32(id), *task, &allTodo)

		fmt.Printf("-- %s is marked as in-progress --\n", task.Title)

	}

	data.SaveDataToFile(&allTodo)

}
