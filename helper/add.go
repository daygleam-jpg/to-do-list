package helper

import (
	"fmt"
	"strings"
	"time"

	"proga/str"

	"github.com/k0kubun/pp"
)

func Add(tasks []str.TooDoo, fields []string) ([]str.TooDoo, string) {
	if len(fields) < 3 {
		errorText := "нужно указать заголовок и текст задачи"
		fmt.Println("Ошибка:", errorText)
		return tasks, errorText
	}

	title := fields[1]
	taskText := strings.Join(fields[2:], " ")

	// Проверяем, есть ли уже задача с таким заголовком.
	for _, task := range tasks {
		if task.Zagolovok == title {
			errorText := "задача с таким заголовком уже существует"
			fmt.Println("Ошибка:", errorText)
			return tasks, errorText
		}
	}

	task := str.TooDoo{
		Zagolovok:   title,
		TextZadachi: taskText,
		CreatedTime: time.Now(),
		Status:      false,
	}

	tasks = append(tasks, task)

	fmt.Println("Вы добавили задачу:")
	pp.Println(task)

	return tasks, ""
}
