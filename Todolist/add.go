package todolist

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func (t *TodoList) Add(task Task) string {
	// Проверяем, есть ли уже задача с таким заголовком.
	for _, existingTask := range t.tasks {
		if existingTask.Title == task.Title {
			errorText := "задача с таким заголовком уже существует"
			fmt.Println("Ошибка:", errorText)
			return errorText
		}
	}

	t.tasks = append(t.tasks, task)

	fmt.Println("Вы добавили задачу:")
	pp.Println(task)

	return ""
}
