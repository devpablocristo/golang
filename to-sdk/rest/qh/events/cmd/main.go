package main

import (
	"log"

	"github.com/devpablocristo/qh/events/cmd/api"
	"github.com/devpablocristo/qh/events/internal/platform/config"
	"go-micro.dev/v4/logger"
)

const (
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorReset  = "\033[0m"
)

// Funciones de log personalizadas
func logInfo(format string, v ...interface{}) {
	log.Printf(ColorBlue+format+ColorReset, v...)
}

func logWarning(format string, v ...interface{}) {
	log.Printf(ColorYellow+format+ColorReset, v...)
}

func logError(format string, v ...interface{}) {
	log.Printf(ColorRed+format+ColorReset, v...)
}

func microLogInfo(format string, v ...interface{}) {
	logger.Infof(ColorBlue+format+ColorReset, v...)
}

func microLogWarning(format string, v ...interface{}) {
	logger.Warnf(ColorYellow+format+ColorReset, v...)
}

func microLogError(format string, v ...interface{}) {
	logger.Errorf(ColorRed+format+ColorReset, v...)
}

// TODO: herramientas consul, kubertes, pprof, gin, viper, k8s, docker, hysterix, jaeger, opentelemetry, wire,

func main() {
	logInfo("Starting application...")
	config, err := config.Load()
	if err != nil {
		logError("error loading config: %v", err)
		return
	}

	router, err := api.InitRouter(config)
	if err != nil {
		microLogError("error initializing router: %v", err)
		return
	}

	if err := api.RunServer(router, config); err != nil {
		microLogError("error running server: %v", err)
	}
}
