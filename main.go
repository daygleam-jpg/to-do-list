package main

import (
	"bufio"
	"fmt"
	"os"
	"proga/str"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tasks := []str.TooDoo{}
	events := []str.Event{}

	for {
		fmt.Print("Введите команду: ")

		if ok := scanner.Scan(); !ok {
			pp.Println("Ошибка ввода")
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

		if cmd == "add" {
			if len(fields) < 3 {
				errorText = "нужно указать заголовок и текст задачи"
				fmt.Println("Ошибка:", errorText)

			} else {
				title := fields[1]
				taskText := strings.Join(fields[2:], " ")

				// Проверяем, есть ли уже задача с таким заголовком
				duplicate := false

				for _, task := range tasks {
					if task.Zagolovok == title {
						duplicate = true
						break
					}
				}

				if duplicate {
					errorText = "задача с таким заголовком уже существует"
					fmt.Println("Ошибка:", errorText)
				} else {
					task := str.TooDoo{
						Zagolovok:   title,
						TextZadachi: taskText,
						CreatedTime: time.Now(),
						Status:      false,
					}

					tasks = append(tasks, task)

					fmt.Println("Вы хотите добавить задачу:")
					pp.Println(task)
				}
			}

		} else if cmd == "list" {
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст")
			} else {
				for i, task := range tasks {
					fmt.Printf(
						"%d. %s — %s\n",
						i+1,
						task.Zagolovok,
						task.TextZadachi,
					)

					fmt.Println("   Создана:", task.CreatedTime)
					fmt.Println("   Выполнена:", task.Status)

					if task.Status {
						fmt.Println("   Время выполнения:", task.MadeTime)
					}
				}
			}

		} else if cmd == "done" {
			if len(fields) < 2 {
				errorText = "укажите заголовок задачи"
				fmt.Println("Ошибка:", errorText)

			} else {
				title := fields[1]
				found := false

				for i := range tasks {
					if tasks[i].Zagolovok == title {
						found = true

						// Время устанавливаем только при первом done
						if !tasks[i].Status {
							tasks[i].Status = true
							tasks[i].MadeTime = time.Now()

							fmt.Println("Задача отмечена как выполненная:", title)
						} else {
							fmt.Println("Задача уже выполнена:", title)
						}

						break
					}
				}

				if !found {
					errorText = "задача не найдена: " + title
					fmt.Println("Задача не найдена:", title)
				}
			}

		} else if cmd == "del" {
			if len(fields) < 2 {
				errorText = "укажите заголовок задачи"
				fmt.Println("Ошибка:", errorText)

			} else {
				title := fields[1]
				found := false

				for i := range tasks {
					if tasks[i].Zagolovok == title {
						tasks = append(tasks[:i], tasks[i+1:]...)

						fmt.Println("Задача удалена:", title)

						found = true
						break
					}
				}

				if !found {
					errorText = "выполненная задача не найдена: " + title
					fmt.Println("Выполненная задача не найдена:", title)
				}
			}

		} else if cmd == "help" {
			fmt.Println("Вы запросили список доступных команд")
			fmt.Println("--------")
			pp.Println("help")
			fmt.Println("Запрашивает список доступных команд")
			fmt.Println("--------")
			pp.Println("add")
			fmt.Println("Добавляет задачу")
			fmt.Println("--------")
			pp.Println("del")
			fmt.Println("Удаляет выполненную задачу")
			fmt.Println("--------")
			pp.Println("list")
			fmt.Println("Запрашивает полный список задач")
			fmt.Println("--------")
			pp.Println("done")
			fmt.Println("Помечает задачу как выполненную")
			fmt.Println("--------")
			pp.Println("events")
			fmt.Println("Запрашивает список всех событий")
			fmt.Println("--------")
			pp.Println("exit")
			fmt.Println("Завершает программу")

		} else if cmd == "events" {
			if len(events) == 0 {
				fmt.Println("Список событий пуст")
				continue
			}

			fmt.Println("Список всех событий:")
			fmt.Println("--------")

			for i, event := range events {
				fmt.Printf("%d. Ввод: %s\n", i+1, event.InputText)

				if event.ErrorText != "" {
					fmt.Println("   Ошибка:", event.ErrorText)
				}

				// Для done показываем первое время выполнения.
				// Для остальных команд показываем время ввода команды.
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

			continue

		} else if cmd == "exit" {
			fmt.Println("Вы завершаете программу")

			event := str.Event{
				InputText:   text,
				ErrorText:   "",
				CreatedTime: time.Now(),
			}

			events = append(events, event)

			return

		} else if cmd == "" {
			errorText = "Вы не ввели команду"
			fmt.Println("Вы не ввели команду")

		} else {
			errorText = "неизвестная команда"
			fmt.Println("Вы ввели неизвестную команду")
		}

		// Создаём событие после выполнения команды.
		event := str.Event{
			InputText:   text,
			ErrorText:   errorText,
			CreatedTime: time.Now(),
		}

		// Если это done, сохраняем время первого выполнения задачи.
		if cmd == "done" && len(fields) >= 2 {
			title := fields[1]

			for i := range tasks {
				if tasks[i].Zagolovok == title {
					event.MadeTime = tasks[i].MadeTime
					break
				}
			}
		}

		// Добавляем событие только один раз.
		events = append(events, event)
	}
}
