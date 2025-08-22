package cmd

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/DimaSU2020/eventscalendar/storage"
)

type Log struct {
	TypeLog  string
	Message  string
	At       time.Time
}

type Logger struct {
	loggerEvents []*Log
	storage      storage.Store
}

func NewLogger(store storage.Store) *Logger{
	return &Logger{
		loggerEvents: make([]*Log,0),
		storage:      store,
	}
}

func (l *Logger) Load() error {
	data, err := l.storage.Load()
	if err != nil {
		return errors.New("Ошибка загрузки из стораджа")
	}
	err = json.Unmarshal(data, &l.loggerEvents)
	return err
}

func (l *Logger) Save() error {
	data, err := json.Marshal(l.loggerEvents)
	if err != nil {
		return errors.New("Ошибка сериализации" + err.Error())
	}
	err = l.storage.Save(data)
	if err != nil {
		return err
	}
	return nil
}

func (l *Logger) NewLog(typeLog string, msg string) *Log {
	log := &Log{
		TypeLog: typeLog,
		Message: msg,
		At: time.Now(),
	}
	return log
}

func (l *Logger) AddLog(typeLog string, msg string) error {
	log := l.NewLog(typeLog, msg)
	l.loggerEvents = append(l.loggerEvents, log)
	return nil
}

func (l *Logger) GetLog() []*Log{
	return l.loggerEvents
}