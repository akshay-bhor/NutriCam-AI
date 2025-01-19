package logger

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

type LogLevels map[string]int

var LogLevelList = LogLevels{
	"debug":    1,
	"info":     2,
	"warn":     3,
	"error":    4,
	"critical": 5,
}

var CurrentLogLevel = LogLevelList[os.Getenv("LOG_LEVEL")]

func Debug(v ...interface{}) {
	if CurrentLogLevel > LogLevelList["debug"] {
		return
	}
	currentTime := time.Now().String()
	fmt.Printf("\n{level: DEBUG, time: %v, error: %v, message: %v}", currentTime, "", v)

}

func Info(v ...interface{}) {
	if CurrentLogLevel > LogLevelList["info"] {
		return
	}
	currentTime := time.Now().String()
	fmt.Printf("\n{level: INFO, time: %v, error: %v, message: %v}", currentTime, "", v)
}

func Warn(err error, v ...interface{}) {
	if CurrentLogLevel > LogLevelList["warn"] {
		return
	}
	currentTime := time.Now().String()
	stack := debug.Stack()
	fmt.Printf("\n{level: WARN, time: %v, error: %v, message: %v, stack: %v}", currentTime, err, v, string(stack))
}

func Error(err error, v ...interface{}) {
	if CurrentLogLevel > LogLevelList["error"] {
		return
	}
	currentTime := time.Now().String()
	stack := debug.Stack()
	fmt.Printf("\n{level: ERROR, time: %v, error: %v, message: %v, stack: %v}", currentTime, err, v, string(stack))
}

func Critical(err error, v ...interface{}) {
	if CurrentLogLevel > LogLevelList["critical"] {
		return
	}
	currentTime := time.Now().String()
	stack := debug.Stack()
	fmt.Printf("\n{level: CRITICAL, time: %v, error: %v, message: %v, stack: %v}", currentTime, err, v, string(stack))
}
