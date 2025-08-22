package main

import (
	"fmt"

	"github.com/DimaSU2020/eventscalendar/calendar"
	"github.com/DimaSU2020/eventscalendar/cmd"
	"github.com/DimaSU2020/eventscalendar/logger"
	"github.com/DimaSU2020/eventscalendar/storage"
)




func main() {
	logger.Init()
	logger.Info("cтарт приложения")

	s := storage.NewJsonStorage("data/json/calendar.json")
	c := calendar.NewCalendar(s)
	// z := storage.NewZipStorage("data/archive/calendar.zip")
	// c := calendar.NewCalendar(z)
	err := c.Load()
	if err != nil {
		fmt.Println("не удалось загрузить json календарь")
	}

	sl := storage.NewJsonStorage("data/log/apphistory.json")
	l := cmd.NewLogger(sl)
	err = l.Load()
	if err != nil {
		logger.Warning("не удалось загрузить json логгер")
		fmt.Println("не удалось загрузить json логгер")
	}

	cli := cmd.NewCmd(c,l)
	cli.Run()
}
