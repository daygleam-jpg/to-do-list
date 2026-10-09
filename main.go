package main

import (
	"bufio"
	"fmt"
	"os"
	"proga/helper"
	"proga/str"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024), 1024*1024)

	tasks := []str.TooDoo{}
	events := []str.Event{}

	for {
		fmt.Print("Введите команду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода")
			return
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
			tasks, errorText = helper.Add(tasks, fields)

		case "list":
			helper.List(tasks)

		case "done":
			tasks, errorText, eventMadeTime = helper.Done(tasks, fields)

		case "del":
			tasks, errorText = helper.Del(tasks, fields)

		case "help":
			helper.Help()

		case "events":
			helper.Events(events)
			continue

		case "exit":
			fmt.Println("Вы завершаете программу")

			event := str.Event{
				InputText:   text,
				ErrorText:   "",
				CreatedTime: time.Now(),
			}

			events = append(events, event)

			return

		default:
			errorText = "неизвестная команда"
			fmt.Println("Вы ввели неизвестную команду")
		}

		// Сохраняем событие после выполнения команды.
		event := str.Event{
			InputText:   text,
			ErrorText:   errorText,
			CreatedTime: time.Now(),
			MadeTime:    eventMadeTime,
		}

		events = append(events, event)
	}
}
