package todolist

import (
	"fmt"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

func Add(tasks []Task, fields []string) ([]Task, string) {
	if len(fields) < 3 {
		errorText := "нужно указать заголовок и текст задачи"
		fmt.Println("Ошибка:", errorText)
		return tasks, errorText
	}

	title := fields[1]
	taskText := strings.Join(fields[2:], " ")

	// Проверяем, есть ли уже задача с таким заголовком.
	for _, task := range tasks {
		if task.Title == title {
			errorText := "задача с таким заголовком уже существует"
			fmt.Println("Ошибка:", errorText)
			return tasks, errorText
		}
	}

	task := Task{
		Title:       title,
		Text:        taskText,
		CreatedTime: time.Now(),
		IsDone:      false,
	}

	tasks = append(tasks, task)

	fmt.Println("Вы добавили задачу:")
	pp.Println(task)

	return tasks, ""
}
