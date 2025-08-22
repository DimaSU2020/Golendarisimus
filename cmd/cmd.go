package cmd

import (
	// "fmt"
	"fmt"
	"sync"

	"github.com/DimaSU2020/eventscalendar/calendar"
	"github.com/c-bata/go-prompt"
)


type Cmd struct {
	calendar *calendar.Calendar
	logger   *Logger
	mu       sync.Mutex
}

func NewCmd(c *calendar.Calendar, l *Logger) *Cmd {
	return &Cmd{
		calendar : c,
		logger   : l,
	}
}

func (c *Cmd) getSuggestions() []prompt.Suggest{
	return []prompt.Suggest{
		{Text: "help", Description: "Показать справку"},
		{Text: "list [ID]", Description: "Показать все события [c ID]"},
		{Text: "add", Description: "Добавить событие"},
		{Text: "remove", Description: "Удалить событие"},
		{Text: "update", Description: "Изменить событие"},
		{Text: "reminder", Description: "Добавить напоминание"},
		{Text: "stop_reminder", Description: "Остановить и удалить напоминание"},
		{Text: "log", Description: "Логирование всех уведомлений (напоминаний), команд и других выводов в консоль"},
		{Text: "exit", Description: "Выйти из программы"},
	}
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest{
	return prompt.FilterHasPrefix(c.getSuggestions(), d.GetWordBeforeCursor(), true)
}

func (c *Cmd) countSuggestions() prompt.Option {
	return prompt.OptionMaxSuggestion(uint16(len(c.getSuggestions())))
}

func (c *Cmd) Run() {
	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
		c.countSuggestions(),
	)
	go func() {
		for msg := range c.calendar.Notification {
			c.mu.Lock()
			fmt.Println(msg)
			typeLog := "message"
			c.logger.AddLog(typeLog, msg)
			c.mu.Unlock()
		}

	}()
	p.Run()
}
