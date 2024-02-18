package common

import (
	"os"
	"laundry-app/config"
	modelutil "laundry-app/utils/model_util"

	"github.com/sirupsen/logrus"
)

type MyLogger interface {
	InitLogger() error // initialize logger
	LogInfo(requestLog modelutil.RequestLog)
	LogWarn(requestLog modelutil.RequestLog)
	LogFatal(requestLog modelutil.RequestLog)
}

type myLogger struct {
	cfg config.LogConfig
	log *logrus.Logger
}

func (m *myLogger) InitLogger() error {
	file, err := os.OpenFile(m.cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	m.log = logrus.New()
	m.log.SetOutput(file)
	return nil
}

func (m *myLogger) LogInfo(requestLog modelutil.RequestLog) {
	m.log.Info(requestLog)
}

func (m *myLogger) LogWarn(requestLog modelutil.RequestLog) {
	m.log.Warn(requestLog)
}

func (m *myLogger) LogFatal(requestLog modelutil.RequestLog) {
	m.log.Fatal(requestLog)
}

func NewMyLogger(cfg config.LogConfig) MyLogger {
	return &myLogger{cfg: cfg}
}
