package logger

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
