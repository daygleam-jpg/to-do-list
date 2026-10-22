package todolist

import "time"

type Task struct {
	Title       string
	Text        string
	CreatedTime time.Time
	IsDone      bool
	MadeTime    time.Time
}

type TodoList struct {
	tasks []Task
}
