package log

import (
	"os"
	"strings"
	"github.com/sirupsen/logrus"
)

type Log struct {
	logrus.Logger
	traceId string
}

var stringToLevel = map[string]logrus.Level{
	"PANIC": logrus.PanicLevel,
	"FATAL": logrus.FatalLevel,
	"ERROR": logrus.ErrorLevel,
	"WARN":  logrus.WarnLevel,
	"INFO":  logrus.InfoLevel,
	"DEBUG": logrus.DebugLevel,
	"TRACE": logrus.TraceLevel,
}

type LogHook struct {
	logger *Log
}

func (h LogHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
}

func (h LogHook) Fire(ent *logrus.Entry) error {
	ent.Data["trace-id"] = h.logger.traceId
	return nil
}

func NewLogger() *Log {
	log := Log{Logger: *logrus.New()}
	log.AddHook(LogHook{logger: &log})
	log.SetLevel(getLogLevelFromEnv())
	file := os.Getenv("LOGFILE")
	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
		if err == nil {
			log.SetOutput(f)
		}
	}
	return &log
}

func (log *Log) SetTraceId(traceId string) {
	log.traceId = traceId
}

func getLogLevelFromEnv() logrus.Level {
	envloglevel := os.Getenv("LOGLEVEL")
	if envloglevel == "" {
		return logrus.InfoLevel
	}
	levelToReturn, ok := stringToLevel[strings.ToUpper(envloglevel)]
	if ok == false {
		levelToReturn = logrus.InfoLevel
	}
	return levelToReturn
}
