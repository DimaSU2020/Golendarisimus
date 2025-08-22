package reminder

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Reminder struct {
	Message string
	At      time.Time
	Sent    bool
	timer   *time.Timer
	Notify  func(string)
}

var	(
	ErrEmptyMessage = errors.New("сообщение пусто")
)

func NewReminder(message string, atTime time.Time, notifier func(string)) (*Reminder, error) {

	if len(strings.TrimSpace(message)) == 0 {
		return nil, fmt.Errorf("нельзя добавить напоминание: %w", ErrEmptyMessage)
	}

	return &Reminder{
		Message : message,
		At      : atTime,
		Sent    : false,
		timer   : nil,
		Notify  : notifier,
	}, nil
}

func (r *Reminder) Start() {
	r.Notify("Стартовали")
	pause := time.Until(r.At)
	r.timer = time.AfterFunc(pause, r.Send)
}

func (r *Reminder) Send() {
	if r.Sent {
		r.Notify("Напоминание уже выслано")
		return
	}
	r.Notify("Напоминание! : "+r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() {
	if r.timer == nil {
		r.Notify("Останавливать нечего, таймер ещё не установлен")
		return
	}
	timer := r.timer
	stopped := timer.Stop()
	if stopped {
		r.Notify("Таймер остановлен до срабатывания")
	} else {
		r.Notify("Таймер уже сработал или был ранее остановлен")
	}
}

