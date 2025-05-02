package logger

import (
	"fmt"
)

// Requirements

/*
1. Support different log levels.
2. Support output to different streams for a given log.
*/

/*
Approach:
 1. Singleton pattern for loggerManager.
 2. Observer pattern for the different kind of logger.
    - FileAppender Logger ->
    - RollingFileAppender -> with log rotation support based on size/time.
    - DB Appender
    - StdOut Logger

3. Thread Safety.
4. Async in nature.
  - Use a que to buffer logs.
  - worker thread consume from queue for actual writing?

5.  Formatter/Layouts
 - Support customizable message formats like:
  [Ts][Ll][Thread] - Message

6. Configuration via a JSON file
 - Dynamically update logging behaviour during runtime.


7. Filters:
 - Allow conditional logging
*/

type Logger struct {
	// name of the logger - need to immutable after creating it.
	name string

	// level is guarded by SetLevel, so callers can’t set nonsense values.
	level LogLevel

	// appender is basically the the list - which can be appended to.
	appenders []Appenders
}

// expose the logger to get inited from outside.
func InitLogger() *Logger {
	return nil
}

// methods on this struct - similar to class methods in java.
func (l *Logger) Debug(message string) {
	// check the log level

}

func (l *Logger) Log(lvl LogLevel, message string) error {
	// 1.  here it needs to check the current level and then log it.
	// 2. lets say the log needs to be sent to all the appenders ...

	if lvl >= l.level {
		// add all the appenders to log it.
		for i := range l.appenders {
			appender := l.appenders[i]
			appender.append(formatMessage(lvl, message))
		}
	}

}

func formatMessage(lvl LogLevel, message string) {
	fmt.Printf()
}
