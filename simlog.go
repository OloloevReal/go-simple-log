package simlog

import "log"

/*
Just simple logs, not more
*/

var version = "0.0.1"

var (
	infoTag  = "[INFO]"
	debugTag = "[DEBUG]"
	errorTag = "[ERROR]"
)

type Logger struct {
	debug bool
}

type Option func(l *Logger)

func SetDebug(l *Logger) {
	l.debug = true
}

func NewLogger(opts ...Option) *Logger {
	l := &Logger{}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *Logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

var globalLog = NewLogger()

func Printf(format string, v ...interface{}) {
	globalLog.Printf(format, v...)
}
