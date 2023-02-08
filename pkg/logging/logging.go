package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err

}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}
func (l *Logger) GetLoggerWithField(s string, v interface{}) *Logger {
	return &Logger{l.WithField(s, v)}
}
func init() {
	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	//err := os.Mkdir("logs", 0644)
	//if err != nil {
	//	panic(err)
	//}
	allfile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.Discard)
	log.AddHook(&writerHook{
		Writer:    []io.Writer{allfile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})
	log.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(log)
}
