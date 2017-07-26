package utils

import (
	"github.com/op/go-logging"
)

var logger *logging.Logger

func Init() error {
	logger = logging.MustGetLogger("case-go")
	return nil
}

func GetLogger() *logging.Logger {
	return logger
}