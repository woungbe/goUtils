package service

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logs struct {
	Path   string
	Logger *logrus.Logger
}

var LogsPTR *Logs

func Init(path string) {
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("로그 파일을 열 수 없습니다:", err)
	}
	defer logFile.Close()

	logger := logrus.New()
	logger.Out = logFile
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	tmplog := new(Logs)
	tmplog.Path = path
	tmplog.Logger = logger
	LogsPTR = tmplog
}

func LoggerError(msg ...interface{}) {
	LogsPTR.Logger.Errorf(msg...)
}

func LoggerInfo(msg ...interface{}) {
	LogsPTR.Logger.Info(msg...)
}
