package logger

import (
	"log"
	"os"
)

// New creates a new logger using stdout
func New() log.Logger {
	return *log.New(os.Stdout, "", 0)
}

// NewError creates a new logger using stderr
func NewError() log.Logger {
	return *log.New(os.Stderr, "", 0)
}

// NewInfo creates a new logger using stdout with timestamps
func NewInfo() log.Logger {
	return *log.New(os.Stdout, "", log.LstdFlags)
}
