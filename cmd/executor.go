package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/DimaSU2020/eventscalendar/events"
	"github.com/DimaSU2020/eventscalendar/logger"
	"github.com/google/shlex"
)

func (c *Cmd) executor(input string) {
	parts, err := shlex.Split(input)
	if err != nil {
		c.calendar.Notify("ошибка парсинга")
	}

	if len(parts) == 0 {return}

	typeLog := "command"
	cmd := strings.ToLower(parts[0])
	c.logger.AddLog(typeLog, cmd)

	switch cmd {
	case "add":
		if len(parts) < 4 {
			c.calendar.Notify("Формат: add \"название события\" \"дата и время\" \"приоритет\"")
			return
		}
		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])

		e, err := c.calendar.AddEvent(title, date, string(priority))
		if err != nil {
			c.calendar.Notify("Ошибка добавления: "+err.Error())
		} else {
			c.calendar.Notify("Событие: \""+e.Title+"\" добавлено")
		}
	case "list":
		option := ""
		if len(parts) > 1 {
			if parts[1] == "ID" || parts[1] == "[ID]" {
				option = "ID"
			} else {
				c.calendar.Notify("Формат: list ID или list [ID]")
				return
			}
		}
		events := c.calendar.GetEvents()
		spaceStr := "  "
		if len(events) == 0 {
			c.calendar.Notify(spaceStr+"*Сожалеем, но у Вас ещё нет событий в календаре")
			return
		}
		c.calendar.Notify("Вот все события Вашего календаря:")
		for _, event := range events {
			c.calendar.Notify(spaceStr+"*Событие: \""+event.Title+
			"\"\n"+spaceStr+" начало: \""+event.StartAt.Format("Monday, 02 Jan 2006, 15:04")+
			"\"\n"+spaceStr+" приоритет: \""+string(event.Priority)+"\"")
			if option == "ID" {
				c.calendar.Notify(spaceStr+" ID: "+event.ID)
			}
			if event.Reminder != nil {
				c.calendar.Notify(spaceStr + " напоминание: \"" + event.Reminder.Message + "\"")
			}
		}
	case "help":
		c.calendar.Notify("Доступные команды:")
		suggestions := c.getSuggestions()
		maxTextLength := 0
		for _, command := range suggestions {
			if len(command.Text) > maxTextLength {
				maxTextLength = len(command.Text)
			}
		}
		for _, command := range suggestions {
			padding := strings.Repeat(" ", maxTextLength-len(command.Text))
			message := fmt.Sprintf("\"%s\"%s: %s", command.Text, padding, command.Description)
			c.calendar.Notify(message)
		}
	case "remove":
		if len(parts) < 2 {
			c.calendar.Notify("Формат: remove \"ID события\"")
			return
		}
		id := parts[1]
		c.calendar.DeleteEvent(id)
	case "update":
		if len(parts) < 4 {
			c.calendar.Notify("Формат: update \"ID события\" \"название события\" \"дата и время\" \"приоритет\"")
			return
		}
		id := parts[1]
		title := parts[2]
		date := parts[3]
		priority := parts[4]
		err := c.calendar.EditEvent(id, title, date, priority)
		if err != nil {
			c.calendar.Notify(err.Error())
		}
	case "reminder":
		if len(parts) < 4 {
			c.calendar.Notify("Формат: reminder \"ID события\"  \"сообщение\" \"время начала или время до события\" ")
			return
		}
		ID := parts[1]
		message := parts[2]
		timeData := parts[3]
		err := c.calendar.SetEventReminder(ID, message, timeData)
		if err != nil {
			c.calendar.Notify("Неудачная попытка добавить напоминание:\n "+err.Error())
		}
	case "stop_reminder":
		if len(parts) < 2 {
			c.calendar.Notify("Формат: stop_reminder \"ID события\"")
			return
		}
		ID := parts[1]
		err := c.calendar.CancelEventReminder(ID)
		if err != nil {
			c.calendar.Notify("Неудачная попытка удалить напоминание:\n "+err.Error())
		}
	case "log":
		logstory := c.logger.GetLog()
		if len(logstory) == 0 {
			fmt.Println("* Сожалеем, но Ваш журнал событий пуст")
			return
		}
		fmt.Println("Вот все записи Вашего журнала событий:")
		for _, log := range logstory {
			fmt.Println("*"+log.At.Format("02.01.06 15:04:00")+", \""+log.TypeLog+"\", \""+log.Message+"\"")
		}
	case "exit":
		c.calendar.Save()
		c.logger.Save()
		c.calendar.CloseCh()
		logger.Notice("запись информации в сторадж и логгер уведомления завершена")
		logger.Info("вышли из приложения")
		os.Exit(0)
	default:
		c.calendar.Notify("Неизвестная команда: "+cmd)
		c.calendar.Notify("Введите 'help' для списка команд")
	}
}
