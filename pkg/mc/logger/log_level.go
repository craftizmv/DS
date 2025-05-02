package logger

import "fmt"

type LogLevel int

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

func printLogLevel() {
	fmt.Println(DEBUG, INFO, WARN)
}
