package main

import (
	"github.com/Hertucktor/archive-api/apiv1"
	"github.com/Hertucktor/archive-api/utils"
)

func main() {
	logFileName := "log_file.json"
	logger := utils.InitializeSugarLogger(logFileName)
	port := ":8080"
	apiv1.SetupRoutes(logger, port)
}
