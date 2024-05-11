package tcplogger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
	"path/filepath"
	"runtime"
)

type writerHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
	Formatter logrus.Formatter
}

func (h *writerHook) Fire(entry *logrus.Entry) error {
	line, err := h.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = h.Writer.Write(line)
	return err
}

func (h *writerHook) Levels() []logrus.Level {
	return h.LogLevels
}

type Logger struct {
	conn net.Conn
}

func NewLogger(level string, logServiceAddr string, setReportCaller bool) (*Logger, error) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.TraceLevel
	}

	logrus.SetReportCaller(setReportCaller)
	logrus.SetLevel(lvl)
	logrus.SetOutput(io.Discard)

	conn, err := net.Dial("tcp", logServiceAddr)
	if err != nil {
		return nil, err
	}

	logstashHook := writerHook{
		Writer:    conn,
		LogLevels: logrus.AllLevels,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				return "", getFileNameAndLine(12)
			},
		},
	}

	stdoutHook := writerHook{
		Writer:    os.Stdout,
		LogLevels: logrus.AllLevels,
		Formatter: &logrus.TextFormatter{
			PadLevelText:           true,
			ForceColors:            true,
			DisableLevelTruncation: true,
			TimestampFormat:        "2006-01-02 15:04:05",
			FullTimestamp:          true,
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				return "", getFileNameAndLine(13)
			},
		},
	}

	logrus.AddHook(&stdoutHook)
	logrus.AddHook(&logstashHook)

	return &Logger{
		conn: conn,
	}, nil
}

func (l *Logger) Close() error {
	return l.conn.Close()
}

func (l *Logger) Trace(message string) {
	logrus.Trace(message)
}

func (l *Logger) Debug(message string) {
	logrus.Debug(message)
}

func (l *Logger) Info(message string) {
	logrus.Info(message)
}

func (l *Logger) Warn(message string) {
	logrus.Warn(message)
}

func (l *Logger) Error(message string) {
	logrus.Error(message)
}

func (l *Logger) Fatal(message string) {
	logrus.Fatal(message)
}

func (l *Logger) Panic(message string) {
	logrus.Panic(message)
}

func getFileNameAndLine(n int) string {
	_, file, line, ok := runtime.Caller(n)
	if !ok {
		return "getting file name and line error"
	}

	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}
