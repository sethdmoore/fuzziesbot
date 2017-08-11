package logconfig

import (
	"github.com/juju/loggo"
)

var log = loggo.GetLogger("")

// New initializes all the loggo loggers
func New(level string) loggo.Logger {
	loggo.ConfigureLoggers(level)
	return log
}

// Get returns the current loggo logger with config
func Get() loggo.Logger {
	return log
}
