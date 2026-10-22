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

	tasks := []todolist.Task{}
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
			tasks, errorText = todolist.Add(tasks, fields)

		case "list":
			todolist.List(tasks)

		case "done":
			tasks, errorText, eventMadeTime = todolist.Done(tasks, fields)

		case "del":
			tasks, errorText = todolist.Del(tasks, fields)

		case "help":
			todolist.Help()

		case "events":
			events.Events(eventsList)

		case "exit":
			fmt.Println("Вы завершаете программу")

			eventsList = append(eventsList, events.NewEvent(text, "", time.Time{}))

			return

		default:
			errorText = "неизвестная команда"
			fmt.Println("Вы ввели неизвестную команду")
		}

		// Сохраняем событие после выполнения команды.
		eventsList = append(eventsList, events.NewEvent(text, errorText, eventMadeTime))
	}

	// Проверяем ошибку scanner после завершения цикла.
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка ввода:", err)
	}
}
