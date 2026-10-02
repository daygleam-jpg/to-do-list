package helper

import (
	"fmt"
	"time"

	"proga/str"
)

func Done(tasks []str.TooDoo, fields []string) ([]str.TooDoo, string, time.Time) {
	if len(fields) < 2 {
		errorText := "укажите хотя бы один заголовок задачи"
		fmt.Println("Ошибка:", errorText)

		return tasks, errorText, time.Time{}
	}

	var eventMadeTime time.Time

	for _, title := range fields[1:] {
		found := false

		for i := range tasks {
			if tasks[i].Zagolovok == title {
				found = true

				if !tasks[i].Status {
					tasks[i].Status = true
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
