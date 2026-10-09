package helper

import (
	"fmt"

	"proga/str"
)

func Events(events []str.Event) {
	if len(events) == 0 {
		fmt.Println("Список событий пуст")
		return
	}

	fmt.Println("Список всех событий:")
	fmt.Println("--------")

	for i, event := range events {
		fmt.Printf(
			"%d. Ввод: %s\n",
			i+1,
			event.InputText,
		)

		if event.ErrorText != "" {
			fmt.Println("   Ошибка:", event.ErrorText)
		}

		if !event.MadeTime.IsZero() {
			fmt.Println(
				"   Время выполнения:",
				event.MadeTime.Format("02.01.2006 15:04:05"),
			)
		} else {
			fmt.Println(
				"   Время:",
				event.CreatedTime.Format("02.01.2006 15:04:05"),
			)
		}

		fmt.Println("--------")
	}
}
