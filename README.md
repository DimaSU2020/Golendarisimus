# 📆 Eventscalendar (Календарь событий)

![Go Version](https://img.shields.io/badge/Go-1.24%2B-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Version](https://img.shields.io/badge/Version-0.1.0beta)

Умный календарь событий с системой напоминаний, реализованный на Go - это приложение, которое позволяет создавать и управлять событиями, устанавливать напоминания и сохранять данные в различных форматах.

## ✨ Возможности

-   ✅ **Добавление событий** с детальной информацией
-   ✏️ **Редактирование и удаление** событий
-   ⏰ **Система напоминаний** с уведомлениями в терминале
-   🗑️ **Удаление напоминаний**
-   💾 **Сохранение данных** в JSON и ZIP форматах
-   📜 **История команд** с просмотром выполненных действий
-   🔍 **Логирование** уведомлений и ошибок
-   🎯 **Интуитивный интерфейс** командной строки

## 🚀 Установка

### Установка через Go get

```bash
go get https://github.com/DimaSU2020/Golendarisimus.git@latest
```

### Сборка из исходного кода

```bash
git clone https://github.com/DimaSU2020/Golendarisimus.git
cd eventscalendar
go build
./eventscalendar
```

📖 Использование, основные команды

```bash
# Показать помощь
help

# Показать все события календаря без ID
list

# Показать все события календаря c ID
list ID

# Добавить новое событие
add "название события" "дата и время события" "приоритет"

# Редактирование события
update "id события" "название события" "дата и время события" "приоритет"

# Удаление события
remove "id события"

# Добавления напоминания для события
reminder "id события" "сообщение для напоминания" "дата и время напоминания"

# Остановка и удаление напоминания для события
cancel_reminder "id события"

# Вывод истории ввода-вывода в консоль
log

# Выход из программы
exit
```

🏗️ Структура проекта

```bash
calendarOfEvents/
├── app.go
├── calendar
│   ├── calendar.go
│   ├── calendarmethods.go
│   └── messages.go
├── cmd
│   ├── cmd.go
│   ├── executor.go
│   ├── executorcomands.go
│   ├── loghistory.go
│   └── messages.go
├── data
│   ├── archive
│   │   └── calendar.zip
│   ├── json
│   │   └── calendar.json
│   └── log
│       ├── app.log
│       └── apphistory.json
├── events
│   ├── event.go
│   ├── messages.go
│   ├── validation_test.go
│   └── validation.go
├── go.mod
├── go.sum
├── LICENSE
├── logger
│   ├── logger.go
│   └── message.go
├── main.go
├── messages.go
├── README.MD
├── reminder
│   ├── messages.go
│   └── reminder.go
└── storage
    ├── json-storage.go
    ├── messages.go
    ├── storage.go
    └── zip-storage.go
```

📋 Примеры использования

```bash
# Добавление события
> add "Тренировка в зале" "2025/08/23 18:00" "high"
# Событие: 'Тренировка в зале' добавлено ID: 'e364996d-03ef-460e-bde1-d39b370a3472'

# Добавление напоминания
> add_reminder "e364996d-03ef-460e-bde1-d39b370a3472" "Напоминаю, сегодня у тебя тренировка в 18:00" "2025/08/23 17:00"
# Напоминание произойдет через 5h29m18.590012s

# Удаление события
> remove "e364996d-03ef-460e-bde1-d39b370a3472"
# Событие: "e364996d-03ef-460e-bde1-d39b370a3472" удалено
```

📊 Форматы данных
Структура события (JSON)

```json
{
    "318400f7-fa52-4989-9b3f-5259cdf1c9bc": {
        "id": "318400f7-fa52-4989-9b3f-5259cdf1c9bc",
        "title": "Тренировка в зале",
        "start_at": "2025-08-23T18:00:00+03:00",
        "Priority": "high",
        "Reminder": null
    }
}
```

Структура события с напоминанием

```json
{
    "318400f7-fa52-4989-9b3f-5259cdf1c9bc": {
        "id":       "318400f7-fa52-4989-9b3f-5259cdf1c9bc",
        "title":    "Тренировка в зале",
        "start_at": "2025-08-23T18:00:00+03:00",
        "Priority": "high",
        "Reminder": {
            "Message": "Напоминаю, сегодня у тебя тренировка в 18:00",
            "At":      "2025-08-23T17:00:00+03:00",
            "Sent":    false
        }
    }
}
```

🛠️ Разработка, требования
Go 1.24.5 или выше

### Запуск в режиме разработки

```bash
git clone https://github.com/DimaSU2020/Golendarisimus.git
cd eventscalendar
go run main.go
```

Запуск тестов

```bash
go test ./...
```

Сборка релиза

```bash
go build
```

## 🤝 Вклад

Мы приветствуем вклад в проект! Для этого:

-   Форкните репозиторий

-   Создайте feature branch (git checkout -b feature/amazing-feature)

-   Закоммитьте изменения (git commit -m 'Add amazing feature')

-   Запушьте branch (git push origin feature/amazing-feature)

-   Откройте Pull Request

## 📄 Лицензия

Этот проект распространяется под лицензией MIT. Подробнее см. в файле LICENSE.

## 🆘 Поддержка

Если у вас возникли вопросы или проблемы:

-   Проверьте Issues

-   Создайте новое Issue с описанием проблемы

-   Укажите версию приложения и шаги для воспроизведения

Project Link: https://github.com/DimaSU2020/Golendarisimus.git

⭐ Не забудьте поставить звезду репозиторию, если проект вам понравился!
