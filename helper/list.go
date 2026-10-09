package helper

import (
	"fmt"
	"proga/str"
)

func List(tasks []str.TooDoo) {
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
			task.Zagolovok,
			task.TextZadachi,
		)

		fmt.Println(
			"   Создана:",
			task.CreatedTime.Format("02.01.2006 15:04:05"),
		)

		fmt.Println("   Выполнена:", task.Status)

		if task.Status {
			fmt.Println(
				"   Время выполнения:",
				task.MadeTime.Format("02.01.2006 15:04:05"),
			)
		}

		fmt.Println("--------")
	}
}
