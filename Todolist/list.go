package todolist

import (
	"fmt"
)

func List(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст")
		return
	}

	fmt.Println("Список задач:")
	fmt.Println("--------")

	for i, task := range tasks {
		fmt.Printf(
			"%d. %s — %s\n",
			i+1,
			task.Title,
			task.Text,
		)

		fmt.Println(
			"   Создана:",
			task.CreatedTime.Format("02.01.2006 15:04:05"),
		)
		fmt.Println("   Выполнена:", task.IsDone)

		if task.IsDone {
			fmt.Println(
				"   Время выполнения:",
				task.MadeTime.Format("02.01.2006 15:04:05"),
			)
		}

		fmt.Println("--------")
	}
}
