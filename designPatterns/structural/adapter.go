package structural

import "fmt"

/*
Adapter is a structural design pattern, which allows incompatible objects to collaborate.
The Adapter acts as a wrapper between two objects. It catches calls for one object and transforms them to format and interface recognizable by the second object.
*/

// Logger is the target interface that your application uses.
type Logger interface {
	LogInfo(message string)
	LogError(message string)
}

// ThirdPartyLogger is the adaptee with an incompatible interface.
type ThirdPartyLogger struct{}

func (l *ThirdPartyLogger) PrintMessage(level string, msg string) {
	fmt.Printf("[%s]: %s\n", level, msg)
}

// LoggerAdapter is the adapter that makes ThirdPartyLogger compatible with Logger.
type LoggerAdapter struct {
	thirdPartyLogger *ThirdPartyLogger
}

func NewLoggerAdapter(thirdPartyLogger *ThirdPartyLogger) *LoggerAdapter {
	return &LoggerAdapter{thirdPartyLogger: thirdPartyLogger}
}

func (a *LoggerAdapter) LogInfo(message string) {
	a.thirdPartyLogger.PrintMessage("INFO", message)
}

func (a *LoggerAdapter) LogError(message string) {
	a.thirdPartyLogger.PrintMessage("ERROR", message)
}

// Caller method
func AdapterCaller() {
	// Create an instance of the third-party logger.
	thirdPartyLogger := &ThirdPartyLogger{}

	// Create an instance of the adapter.
	logger := NewLoggerAdapter(thirdPartyLogger)

	// Use the logger interface.
	logger.LogInfo("This is an info message.")
	logger.LogError("This is an error message.")
}
