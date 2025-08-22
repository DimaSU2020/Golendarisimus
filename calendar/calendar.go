package calendar

import (
	"encoding/json"
	"errors"

	"github.com/DimaSU2020/eventscalendar/events"
	"github.com/DimaSU2020/eventscalendar/logger"
	"github.com/DimaSU2020/eventscalendar/storage"
)

type Calendar struct {
	calendarEvents  map[string]*events.Event
	storage         storage.Store
	Notification    chan string
}

func NewCalendar(store storage.Store) *Calendar {
	logger.Notice("создан календарь")
	logger.Notice("создан канал c.calendar.Notification")
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
		storage:        store,
		Notification:   make(chan string),
	}
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return errors.New("Ошибка загрузки из стораджа")
	}
	err = json.Unmarshal(data, &c.calendarEvents)
	return err
}

func (c *Calendar) Save() error {
	data, err := json.Marshal(c.calendarEvents)
	if err != nil {
		return errors.New("Ошибка сериализации" + err.Error())
	}
	err = c.storage.Save(data)
	if err != nil {
		return err
	}
	return nil
}
