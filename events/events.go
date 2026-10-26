package events

import (
	"fmt"
	"time"
)

type Event struct {
	InputText   string
	ErrorText   string
	CreatedTime time.Time
	MadeTime    time.Time
}

func NewEvent(inputText, errorText string, madeTime time.Time) Event {
	return Event{
		InputText:   inputText,
		ErrorText:   errorText,
		CreatedTime: time.Now(),
		MadeTime:    madeTime,
	}
}
func (e Event) Print() {
	fmt.Printf("Команда: %s\n", e.InputText)
	fmt.Printf("Время: %s\n", e.CreatedTime.Format("02.01.2006 15:04:05"))

	if e.ErrorText != "" {
		fmt.Printf("Ошибка: %s\n", e.ErrorText)
	}

	if !e.MadeTime.IsZero() {
		fmt.Printf("Время выполнения: %s\n", e.MadeTime.Format("02.01.2006 15:04:05"))
	}
}
func PrintEvents(eventsList []Event) {
	if len(eventsList) == 0 {
		fmt.Println("Список событий пуст")
		return
	}

	fmt.Println("Список всех событий:")
	fmt.Println("--------")

	for i, event := range eventsList {
		fmt.Printf("%d. ", i+1)
		event.Print()
		fmt.Println("--------")
	}
}
