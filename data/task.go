package data

type Task struct {
	Id     uint32 `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func CreateTask(title string, id uint32) *Task {
	return &Task{Id: id, Title: title, Status: "todo"}

}
