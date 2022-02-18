package config

import (
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

var _isLoggerReady = loggerSetup()

func loggerSetup() bool {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)

			return "", fileName
		},
	})

	logrus.Info("Created logger")

	return true
}

func IsLoggerReady() bool {
	return _isLoggerReady
}
