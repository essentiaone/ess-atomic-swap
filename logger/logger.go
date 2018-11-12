package logger

import (
	"io"
	"os"

	"github.com/essentiaone/ess-atomic-swap/common"
	"github.com/sirupsen/logrus"
)

const (
	timeTemplate = "Monday 02 Jan 15:04:05 MST 2006"
)

// Log create logger
type Log struct {
	Logger *logrus.Logger
}

// New create new logger
func New(isProduction bool) (*Log, error) {
	std, err := chooseOutput(isProduction)
	if err != nil {
		return nil, err
	}

	fieldMap := logrus.FieldMap{
		logrus.FieldKeyLevel: "s",
		logrus.FieldKeyTime:  "t",
		logrus.FieldKeyMsg:   "m",
	}

	textTemplate := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: timeTemplate,
		FieldMap:        fieldMap,
	}

	log := logrus.New()
	log.SetOutput(std)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(textTemplate)

	return &Log{Logger: log}, nil
}

func chooseOutput(isProduction bool) (io.Writer, error) {
	if isProduction {
		return os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	}

	return os.Stderr, nil
}

// Log write 'message' with 'logType' to choose stdout
func (l *Log) Log(logType int, messages ...interface{}) {
	switch logType {
	case common.Error:
		l.Logger.Error(messages)
	case common.Fatal:
		l.Logger.Fatal(messages)
	case common.Info:
		l.Logger.Info(messages)
	case common.Panic:
		l.Logger.Panic(messages)
	case common.Warn:
		l.Logger.Panic(messages)
	default:
		l.Logger.Error(messages)
	}
}
