package helper

import "fmt"

func Help() {
	fmt.Println("Вы запросили список доступных команд")
	fmt.Println("--------")

	fmt.Println("help")
	fmt.Println("Запрашивает список доступных команд")
	fmt.Println("Формат: help")
	fmt.Println("--------")

	fmt.Println("add")
	fmt.Println("Добавляет задачу")
	fmt.Println("Формат: add <Заголовок> <Текст>")
	fmt.Println("--------")

	fmt.Println("del")
	fmt.Println("Удаляет задачу")
	fmt.Println("Формат: del <Заголовок>")
	fmt.Println("--------")

	fmt.Println("list")
	fmt.Println("Запрашивает полный список задач")
	fmt.Println("Формат: list")
	fmt.Println("--------")

	fmt.Println("done")
	fmt.Println("Помечает задачу как выполненную")
	fmt.Println("Формат: done <Заголовок>")
	fmt.Println("--------")

	fmt.Println("events")
	fmt.Println("Запрашивает список всех событий")
	fmt.Println("Формат: events")
	fmt.Println("--------")

	fmt.Println("exit")
	fmt.Println("Завершает программу")
	fmt.Println("Формат: exit")
	fmt.Println("--------")
}
