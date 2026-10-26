package todolist

import (
	"fmt"
	"time"
)

func (t *TodoList) Done(title string) (string, time.Time) {
	for i := range t.tasks {
		if t.tasks[i].Title == title {
			if t.tasks[i].IsDone {
				fmt.Println("Задача уже выполнена:", title)
				return "", time.Time{}
			}

			t.tasks[i].IsDone = true
			t.tasks[i].MadeTime = time.Now()

			fmt.Println("Задача отмечена как выполненная:", title)

			return "", t.tasks[i].MadeTime
		}
	}

	errorText := "задача не найдена: " + title
	fmt.Println("Ошибка:", errorText)

	return errorText, time.Time{}
}
