package calendar

import (
	"errors"
	"fmt"

	"github.com/DimaSU2020/eventscalendar/events"
	"github.com/DimaSU2020/eventscalendar/logger"
)

func (c *Calendar) AddEvent(title, date, priority string) (*events.Event, error) {
	event, err := events.NewEvent(title, date, priority)
	if err != nil {
		return nil, errors.New("Ошибка создания ивента")
	}
	c.calendarEvents[event.ID] = event
	return event, nil
}

func (c *Calendar) DeleteEvent(id string) {
	_, exist := c.calendarEvents[id]
	if !exist {
		c.Notify("не удалось удалить задачу, ключ \""+id+"\" не найден")
		return
	}
	c.Notify("Событие \""+c.calendarEvents[id].Title+"\" успешно удалено")
	delete(c.calendarEvents, id)
}

func (c *Calendar) EditEvent(id, title, date, priority string) error {
    event, exist := c.calendarEvents[id]
    if !exist {
        return errors.New("не удалось отредактировать событие, ключ \""+id+"\" не найден")
    }

    err := event.Update(title, date, priority)
	if err != nil {
		return errors.New("Произошла ошибка: "+err.Error())
	} else {
		c.Notify("Событие была успешно изменено")
	}
    return err
}

var ErrReminderKeyNotFound = errors.New("не удалось установить напоминание, ключ не найден")
var ErrAddReminder = errors.New("не удалось установить напоминание")

func (c *Calendar) SetEventReminder(eventId string, message string, timeData string) error {
	event, exist := c.calendarEvents[eventId]
    if !exist {
        return ErrReminderKeyNotFound
    }
	err := event.AddReminder(message, timeData, c.Notify)
	if err != nil {
		return fmt.Errorf("ошибка установки напоминания: %w", ErrAddReminder)
	}
	return nil
}

func (c *Calendar) CancelEventReminder(eventId string) error {
	event, exist := c.calendarEvents[eventId]
    if !exist {
        return errors.New("не удалось установить напоминание событию, ключ \""+eventId+"\" не найден")
    }
	event.RemoveReminder()
	return nil
}

func (c *Calendar) Notify(msg string) {
	logger.Notice("отправляем сообщение в канал c.calendar.Notification уведомлений")
	c.Notification <- msg
}

func (c*Calendar) CloseCh() {
	logger.Notice("закрываем канал c.calendar.Notification уведомлений")
	close(c.Notification)
}

func (c *Calendar) GetEvents() map[string]*events.Event{
	return c.calendarEvents
}
