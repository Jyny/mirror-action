package logger

import (
	"bytes"
	"io"
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

const (
	// DebugLevel is the debug level.
	DebugLevel = log.DebugLevel
	// InfoLevel is the info level.
	InfoLevel = log.InfoLevel
	// WarnLevel is the warn level.
	WarnLevel = log.WarnLevel
	// ErrorLevel is the error level.
	ErrorLevel = log.ErrorLevel
	// FatalLevel is the fatal level.
	FatalLevel = log.FatalLevel
)

type Logger interface {
	io.Writer
	SetLevel(level log.Level)
	Debug(msg interface{}, keyvals ...interface{})
	Info(msg interface{}, keyvals ...interface{})
	Warn(msg interface{}, keyvals ...interface{})
	Error(msg interface{}, keyvals ...interface{})
	Fatal(msg interface{}, keyvals ...interface{})
	Print(msg interface{}, keyvals ...interface{})
}
type LoggerWthWriter struct {
	io.Writer
	*log.Logger
	lock sync.Mutex
}

func New(writer io.Writer, level log.Level) *LoggerWthWriter {
	logger := log.NewWithOptions(writer, log.Options{
		ReportTimestamp: true,
		Level:           level,
	})
	return &LoggerWthWriter{
		writer,
		logger,
		sync.Mutex{},
	}
}

func (l *LoggerWthWriter) Write(p []byte) (n int, err error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	buf := []byte{}
	for _, b := range bytes.Split(p, []byte("\r")) {
		buf = append(buf, []byte(time.Now().Format("2006-01-02 15:04:05 "))...)
		buf = append(buf, b...)
		buf = append(buf, '\r')
	}

	return l.Writer.Write(buf)
}
