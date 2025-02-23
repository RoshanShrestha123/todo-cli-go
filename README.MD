### Description

Todo CLI application on go (without external tools for CLI like cobra/viper)

__This is the practice project given by the go Roadmap.sh__

https://roadmap.sh/projects/task-tracker

Here’s the information converted into a markdown table:

| Command                                   | Description                           |
|-------------------------------------------|---------------------------------------|
| `task-cli add "Buy groceries"`           | Task added successfully (ID: 1)      |
| `task-cli update 1 "Buy groceries and cook dinner"` | Update task with ID 1               |
| `task-cli delete 1`                       | Delete task with ID 1                |
| `task-cli mark-in-progress 1`            | Mark task with ID 1 as in progress   |
| `task-cli mark-done 1`                   | Mark task with ID 1 as done          |
| `task-cli list`                           | List all tasks                       |
| `task-cli list done`                     | List tasks that are done             |
| `task-cli list todo`                     | List tasks that are todo             |
| `task-cli list in-progress`              | List tasks that are in progress      |


### Features

- [x] Add, Update, and Delete tasks  
- [x] Mark a task as in progress or done  
- [x] List all tasks  
- [x] List all tasks that are done  
- [x] List all tasks that are not done  
- [x] List all tasks that are in progress  

