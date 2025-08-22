package events

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

func (p Priority) ToLower() Priority{
	priority := string(p)
	priority = strings.ToLower(priority)
	return Priority(priority)
}

func CheckingData(title, dateStr, priority string) (bool, time.Time, bool) {
	matched, err := IsValidTitle(title)
	if !matched {
		fmt.Println("Ошибка: неверный формат заголовка")
		return false, time.Time{}, false
	} else if err != nil {
		fmt.Println("Ошибка:", err)
		return false, time.Time{}, false
	}
	timeEvent, err := IsValidDate(dateStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return false, time.Time{}, false
	}
	p := Priority(priority)
	err = IsValidPriority(p)
	if err != nil {
		return false, time.Time{}, false
	}
	return matched, timeEvent, true
}

func IsValidPriority(p Priority) error {
	p = p.ToLower()
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return nil
	default:
		return errors.New("ошибка при задании приоритета")
	}
}

func IsValidTitle(title string) (bool, error) {
	pattern := `^[a-zA-ZА-Яа-я0-9 ,.#]{3,100}$`
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false, errors.New("неверный шаблон или другая ошибка")
	}
	return matched, nil
}

func IsValidDate(dataStr string) (time.Time, error) {
	timeEvent, err := dateparse.ParseLocal(dataStr)
	if err != nil {
		return time.Time{}, errors.New("неверный формат даты")
	}
	return timeEvent, err
}
