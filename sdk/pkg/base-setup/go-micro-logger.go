package BasicSetup

import (
	"fmt"
	"log"
	"os"

	"go-micro.dev/v4/logger"
)

const (
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorReset  = "\033[0m"
)

func LogInfo(format string, v ...interface{}) {
	log.Printf(ColorBlue+format+ColorReset, v...)
}

func LogWarning(format string, v ...interface{}) {
	log.Printf(ColorYellow+format+ColorReset, v...)
}

func LogError(format string, v ...interface{}) {
	log.Printf(ColorRed+format+ColorReset, v...)
}

func MicroLogInfo(format string, v ...interface{}) {
	logger.Infof(ColorBlue+format+ColorReset, v...)
}

func MicroLogWarning(format string, v ...interface{}) {
	logger.Warnf(ColorYellow+format+ColorReset, v...)
}

func MicroLogError(format string, v ...interface{}) {
	logger.Errorf(ColorRed+format+ColorReset, v...)
}

func SetupLogger() error {
	err := logger.Init(
		logger.WithLevel(logger.InfoLevel),
		logger.WithOutput(os.Stdout),
	)
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}
	return nil
}
