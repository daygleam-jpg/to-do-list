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

func NewTask(title, text string) Task {
	return Task{
		Title:       title,
		Text:        text,
		CreatedTime: time.Now(),
		IsDone:      false,
	}
}

func NewTodoList() *TodoList {
	return &TodoList{
		tasks: []Task{},
	}
}
