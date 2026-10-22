package todolist

import (
	"fmt"
	"time"
)

func Done(tasks []Task, fields []string) ([]Task, string, time.Time) {
	if len(fields) < 2 {
		errorText := "укажите хотя бы один заголовок задачи"
		fmt.Println("Ошибка:", errorText)

		return tasks, errorText, time.Time{}
	}

	var eventMadeTime time.Time

	for _, title := range fields[1:] {
		found := false

		for i := range tasks {
			if tasks[i].Title == title {
				found = true

				if !tasks[i].IsDone {
					tasks[i].IsDone = true
					tasks[i].MadeTime = time.Now()

					fmt.Println(
						"Задача отмечена как выполненная:",
						title,
					)

					// Сохраняем время первой выполненной задачи.
					if eventMadeTime.IsZero() {
						eventMadeTime = tasks[i].MadeTime
					}
				} else {
					fmt.Println(
						"Задача уже выполнена:",
						title,
					)
				}

				break
			}
		}

		if !found {
			fmt.Println("Задача не найдена:", title)
		}
	}

	return tasks, "", eventMadeTime
}
