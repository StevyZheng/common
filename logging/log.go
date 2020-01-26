package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = NewLogx()

func FatalPrintln(msg string) {
	if Logger.logger.Out != os.Stdout {
		println(msg)
	}
	Logger.LogFatal(msg)
}

type LogLevel logrus.Level

const (
	InfoLevel  logrus.Level = logrus.InfoLevel
	WarnLevel  logrus.Level = logrus.WarnLevel
	ErrorLevel logrus.Level = logrus.ErrorLevel
	FatalLevel logrus.Level = logrus.FatalLevel
	DebugLevel logrus.Level = logrus.DebugLevel
)

type Logx struct {
	logger *logrus.Logger
}

func NewLogx() *Logx {
	logx := new(Logx)
	logx.logger = logrus.New()
	logx.logger.Out = os.Stdout
	logx.logger.Formatter = &logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    true,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
	}
	logx.logger.Level = logrus.InfoLevel
	return logx
}

func (l *Logx) SetLogLevel(level LogLevel) {
	l.logger.SetLevel((logrus.Level)(level))
}

func (l Logx) LogInfo(msg string) {
	l.logger.WithFields(logrus.Fields{}).Info(msg)
}

func (l Logx) LogFatal(msg string) {
	l.logger.WithFields(logrus.Fields{}).Fatal(msg)
}

func (l Logx) LogWarn(msg string) {
	l.logger.WithFields(logrus.Fields{}).Warn(msg)
}

func (l Logx) LogDebug(msg string) {
	l.logger.WithFields(logrus.Fields{}).Debug(msg)
}
func (l Logx) LogError(msg string) {
	l.logger.WithFields(logrus.Fields{}).Error(msg)
}
