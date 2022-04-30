package log

import "github.com/phuslu/log"

const (
	MaxStackSize = 4096
)

var logger *log.Logger

func init() {
	initLogger()
}

func initLogger() {
	logger = &log.Logger{
		Level:      log.TraceLevel,
		TimeField:  "date",
		TimeFormat: "2006-01-02 12:21:56",
		Writer: &log.ConsoleWriter{
			ColorOutput: true,
			QuoteString: true,
			Writer: &log.FileWriter{
				Filename:     "log/main.log",
				FileMode:     0600,
				MaxSize:      100 * 1024 * 1024,
				EnsureFolder: true,
				LocalTime:    true,
			},
		},
	}
}

func GetLogger() *log.Logger {
	if logger == nil {
		initLogger()
	}
	return logger
}
