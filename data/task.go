package data

type Task struct {
	Id          uint32 `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

func CreateTask(title string, len int) *Task {
	return &Task{Id: uint32(len), Title: title, IsCompleted: false}

}
