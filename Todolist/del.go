package todolist

import (
	"fmt"
)

func Del(tasks []Task, fields []string) ([]Task, string) {
	if len(fields) < 2 {
		errorText := "укажите хотя бы один заголовок задачи"
		fmt.Println("Ошибка:", errorText)

		return tasks, errorText
	}

	for _, title := range fields[1:] {
		found := false

		for i := range tasks {
			if tasks[i].Title == title {
				tasks = append(tasks[:i], tasks[i+1:]...)

				fmt.Println("Задача удалена:", title)

				found = true
				break
			}
		}

		if !found {
			fmt.Println("Задача не найдена:", title)
		}
	}

	return tasks, ""
}
