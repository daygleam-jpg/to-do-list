package str

import "time"

type TooDoo struct {
	Zagolovok   string
	TextZadachi string
	CreatedTime time.Time
	Status      bool
	MadeTime    time.Time
}

type Event struct {
	InputText   string
	ErrorText   string
	CreatedTime time.Time
}
