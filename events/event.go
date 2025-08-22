package events

import (
	"errors"
	"time"
	"github.com/DimaSU2020/eventscalendar/reminder"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID        string              `json:"id"`
	Title     string              `json:"title"`
	StartAt   time.Time           `json:"start_at"`
	Priority  Priority            `json:"priority"`
	Reminder  *reminder.Reminder  `json:"reminder"`
}

func NewEvent(title, dateStart, priority string) (*Event, error){
	matched, timeEvent, yesno := CheckingData(title, dateStart, priority)
	if !matched || !yesno {
		return nil, errors.New("Не удалось добавить событие из-за ошибки в данных")
	}

	return &Event{
		ID: getNextID(),
		Title: title,
		StartAt: timeEvent,
		Priority: Priority(priority),
		Reminder: nil,
	}, nil
}

func getNextID() string{
	return uuid.New().String()
}

func (event *Event) Update(title string, dateStart string, priority string) error {
	matched, timeEvent, yesno := CheckingData(title, dateStart, priority)
	if !matched || !yesno {
		return errors.New("Не удалось изменить событие из-за ошибки в данных")
	}
	event.Title = title
	event.StartAt = timeEvent
	event.Priority = Priority(priority)
	return nil
}

var ErrTooLateEvent    = errors.New("время начала события в прошлом")
var ErrTooLateReminder = errors.New("время напоминания позже события")
var ErrParseTime       = errors.New("ошибка при парсинге данных по времени начала напоминания")


func (e *Event) AddReminder(message string, timeData string, notifier func(string)) error {
	now := time.Now().In(e.StartAt.Location())
	if e.StartAt.Before(now) {
		return ErrTooLateEvent
	}
	date, err := dateparse.ParseLocal(timeData)
    if err != nil {
        duration, err := time.ParseDuration(timeData)
		if err != nil {
			return ErrParseTime
		}
		date = e.StartAt.Add(-duration)
    }
	if date.Before(now) {
		return ErrTooLateReminder
	}
	r, err := reminder.NewReminder(message, date, notifier)
	if err != nil {
		return err
	}
	e.Reminder = r
	e.Reminder.Start()
	return nil
}

func (e *Event) RemoveReminder() {
	if e.Reminder != nil {
		e.Reminder.Stop()
		e.Reminder = nil
	}
}
