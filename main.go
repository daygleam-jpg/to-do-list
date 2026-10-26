package main

import (
	"bufio"
	"fmt"
	"os"
	todolist "proga/Todolist"
	"proga/events"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024), 1024*1024)

	todoList := todolist.NewTodoList()
	eventsList := []events.Event{}

	for {
		fmt.Print("Введите команду: ")

		if !scanner.Scan() {
			break
		}

		text := scanner.Text()
		fields := strings.Fields(text)

		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели")
			continue
		}

		cmd := strings.ToLower(fields[0])
		errorText := ""
		var eventMadeTime time.Time

		switch cmd {
		case "add":
			if len(fields) < 3 {
				errorText = "нужно указать заголовок и текст задачи"
				fmt.Println("Ошибка:", errorText)
				break
			}

			title := fields[1]
			taskText := strings.Join(fields[2:], " ")

			task := todolist.NewTask(title, taskText)
			errorText = todoList.Add(task)

		case "list":
			todoList.List()

		case "done":
			if len(fields) < 2 {
				errorText = "укажите хотя бы один заголовок задачи"
				fmt.Println("Ошибка:", errorText)
				break
			}

			for _, title := range fields[1:] {
				err, madeTime := todoList.Done(title)

				if err != "" {
					errorText = err
				}

				if eventMadeTime.IsZero() && !madeTime.IsZero() {
					eventMadeTime = madeTime
				}
			}

		case "del":
			if len(fields) < 2 {
				errorText = "укажите хотя бы один заголовок задачи"
				fmt.Println("Ошибка:", errorText)
				break
			}

			for _, title := range fields[1:] {
				err := todoList.Del(title)

				if err != "" {
					errorText = err
				}
			}

		case "help":
			todolist.Help()

		case "events":
			events.PrintEvents(eventsList)

		case "exit":
			fmt.Println("Вы завершаете программу")
			eventsList = append(
				eventsList,
				events.NewEvent(text, "", time.Time{}),
			)
			return

		default:
			errorText = "неизвестная команда"
			fmt.Println("Вы ввели неизвестную команду")
		}

		// Сохраняем событие после выполнения команды.
		eventsList = append(
			eventsList,
			events.NewEvent(text, errorText, eventMadeTime),
		)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка ввода:", err)
	}
}
