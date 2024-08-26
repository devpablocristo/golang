package sdklogger

import (
	"log"

	"go-micro.dev/v4/logger"
)

const (
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorReset  = "\033[0m"
)

func StdInfo(format string, v ...any) {
	log.Printf(ColorBlue+format+ColorReset, v...)
}

func StdWarning(format string, v ...any) {
	log.Printf(ColorYellow+format+ColorReset, v...)
}

func StdError(format string, v ...any) {
	log.Printf(ColorRed+format+ColorReset, v...)
}

func GmInfo(format string, v ...any) {
	logger.Infof(ColorBlue+format+ColorReset, v...)
}

func GmWarning(format string, v ...any) {
	logger.Warnf(ColorYellow+format+ColorReset, v...)
}

func GmError(format string, v ...any) {
	logger.Errorf(ColorRed+format+ColorReset, v...)
}
