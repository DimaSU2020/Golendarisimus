package logger

import (
	"log"
	"os"
	"sync"
)

type Logger struct {
	infoLogger     *log.Logger
	noticeLogger   *log.Logger
	warningLogger  *log.Logger
	errorLogger    *log.Logger
}

const (
	PrefixInfo    string = "INFO   : "
	PrefixNotice  string = "NOTICE : "
	PrefixWarning string = "WARNING: "
	PrefixError   string = "ERROR  : "
)

var (
	logger                 *Logger
	once                   sync.Once
	logFilePath            = "data/log/app.log"
	ErrOpenFileLog         = "Ошибка открытия файла логов: "
	ErrFatalInitialisation = "Логгер не инициализирован. Вызовите logger.Init() в main.go"
)

func Init()  {
	once.Do(func() {
		file,err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(ErrOpenFileLog,err)
		}
		log.SetOutput(file)
		logger = &Logger{
			infoLogger     : log.New(file, PrefixInfo,    log.Ldate|log.Ltime|log.Lshortfile),
			noticeLogger   : log.New(file, PrefixNotice,  log.Ldate|log.Ltime|log.Lshortfile),
			warningLogger  : log.New(file, PrefixWarning, log.Ldate|log.Ltime|log.Lshortfile),
			errorLogger    : log.New(file, PrefixError,   log.Ldate|log.Ltime|log.Lshortfile),
		}
	})
}

func getLogger() *Logger{
	if logger == nil {
		log.Fatal(ErrFatalInitialisation)
	}
	return logger
}

func Info(msg string) {
	getLogger().infoLogger.Output(2, msg)
}

func Notice(msg string) {
	getLogger().noticeLogger.Output(2, msg)
}

func Warning(msg string) {
	getLogger().warningLogger.Output(2, msg)
}

func Error(msg string) {
	getLogger().errorLogger.Output(2, msg)
}
