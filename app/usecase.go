package app

import "logger_demo/logger"

func Something1() {
	// Do something
	logger.Info("message")
}

func Something2() {
	// Do something
	logger.Error("message")
}
