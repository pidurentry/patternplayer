package main

import (
	"fmt"
	"os"

	"github.com/pidurentry/buttplug-go/logging"
)

func init() {
	logging.SetLogger(&Logger{DEBUG})
}

type level int

const (
	TRACE = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	level level
}

func (logger *Logger) log(level level, format string, args ...interface{}) {
	if logger.level > level {
		return
	}

	switch level {
	case TRACE:
		fmt.Printf("TRACE: "+format+"\n", args...)
	case DEBUG:
		fmt.Printf("DEBUG: "+format+"\n", args...)
	case INFO:
		fmt.Printf(" INFO: "+format+"\n", args...)
	case WARNING:
		fmt.Printf(" WARN: "+format+"\n", args...)
	case ERROR:
		fmt.Printf("ERROR: "+format+"\n", args...)
	case FATAL:
		fmt.Printf("FATAL: "+format+"\n", args...)
		os.Exit(1)
	}
}

func (logger *Logger) Trace(message interface{}) {
	logger.log(TRACE, "%s", message)
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	logger.log(TRACE, format, args...)
}

func (logger *Logger) Debug(message interface{}) {
	logger.log(DEBUG, "%s", message)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.log(DEBUG, format, args...)
}

func (logger *Logger) Info(message interface{}) {
	logger.log(INFO, "%s", message)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.log(INFO, format, args...)
}

func (logger *Logger) Warning(message interface{}) {
	logger.log(WARNING, "%s", message)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	logger.log(WARNING, format, args...)
}

func (logger *Logger) Error(message interface{}) {
	logger.log(ERROR, "%s", message)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.log(ERROR, format, args...)
}

func (logger *Logger) Fatal(message interface{}) {
	logger.log(FATAL, "%s", message)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	defer os.Exit(1)
	logger.log(FATAL, format, args...)
}
