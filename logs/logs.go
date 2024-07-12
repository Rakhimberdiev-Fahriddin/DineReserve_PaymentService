package logs

import (
	"log"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() *slog.Logger {
	LogFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	handler := slog.NewJSONHandler(LogFile, nil)
	Logger = slog.New(handler)
	return Logger
}
