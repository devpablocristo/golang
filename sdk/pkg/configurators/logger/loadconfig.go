package sdklogger

import (
	"log"

	"go-micro.dev/v4/logger"
)

const (
	red    = "\033[31m"
	yellow = "\033[33m"
	green  = "\033[32m"
	blue   = "\033[34m"
	reset  = "\033[0m"
)

// applyColor applies a color to a string format.
func applyColor(color, format string) string {
	return color + format + reset
}

// Standard log functions using the default log package
func Info(format string, v ...any) {
	log.Printf(applyColor(blue, format), v...)
}

func Warning(format string, v ...any) {
	log.Printf(applyColor(yellow, format), v...)
}

func Error(format string, v ...any) {
	log.Printf(applyColor(red, format), v...)
}

// Go-Micro log functions using the go-micro logger
func Minfo(format string, v ...any) {
	logger.Infof(applyColor(blue, format), v...)
}

func Mwarning(format string, v ...any) {
	logger.Warnf(applyColor(yellow, format), v...)
}

func Merror(format string, v ...any) {
	logger.Errorf(applyColor(red, format), v...)
}
