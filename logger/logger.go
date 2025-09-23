package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

// Logger encapsulates a zerolog.Logger
type Logger struct {
	logger zerolog.Logger
}

var (
	globalLogger *Logger
	mu           sync.Mutex
)

// GetLogger returns the logger instance with the specified verbosity
func GetLogger(verbosity int) *Logger {
	mu.Lock()
	defer mu.Unlock()

	// Configure ConsoleWriter with custom formatting
	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			var color int
			var level string
			if ll, ok := i.(string); ok {
				switch ll {
				case "debug":
					color = 36 // Cyan
					level = "DEBUG"
				case "info":
					color = 32 // Green
					level = "INFO"
				case "warn":
					color = 33 // Yellow
					level = "WARN"
				case "error":
					color = 31 // Red
					level = "ERROR"
				default:
					color = 37
					level = fmt.Sprintf("%v", i)
				}
			}
			return fmt.Sprintf("\x1b[%dm%-6s\x1b[0m", color, level)
		},
		FormatTimestamp: func(i interface{}) string {
			return fmt.Sprintf("\x1b[90m%s\x1b[0m", i)
		},
		FormatFieldName: func(i interface{}) string {
			return fmt.Sprintf("\x1b[94m%s=\x1b[0m", i)
		},
		FormatFieldValue: func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		},
	}

	// Map verbosity to zerolog level
	level, exists := map[int]zerolog.Level{
		0: zerolog.ErrorLevel,
		1: zerolog.WarnLevel,
		2: zerolog.InfoLevel,
		3: zerolog.DebugLevel,
	}[verbosity]
	if !exists {
		level = zerolog.ErrorLevel
	}

	zlogger := zerolog.New(writer).With().Timestamp().Logger().Level(level)
	globalLogger = &Logger{logger: zlogger}
	return globalLogger
}

// Debug logs a message at Debug level
func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debug().Fields(toFields(keysAndValues)).Msg(msg)
}

// Info logs a message at Info level
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info().Fields(toFields(keysAndValues)).Msg(msg)
}

// Warn logs a message at Warn level
func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warn().Fields(toFields(keysAndValues)).Msg(msg)
}

// Error logs a message at Error level
func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Error().Fields(toFields(keysAndValues)).Msg(msg)
}

// toFields converts variadic key-value pairs to a map for zerolog
func toFields(keysAndValues []interface{}) map[string]interface{} {
	if len(keysAndValues)%2 != 0 {
		return nil
	}
	fields := make(map[string]interface{})
	for i := 0; i < len(keysAndValues); i += 2 {
		key, ok := keysAndValues[i].(string)
		if !ok {
			continue
		}
		fields[key] = keysAndValues[i+1]
	}
	return fields
}
