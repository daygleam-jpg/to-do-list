package todolist

import "fmt"

func (t *TodoList) Del(title string) string {
	for i := range t.tasks {
		if t.tasks[i].Title == title {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			fmt.Println("Задача удалена:", title)
			return ""
		}
	}

	errorText := "задача не найдена: " + title
	fmt.Println("Ошибка:", errorText)
	return errorText
}
