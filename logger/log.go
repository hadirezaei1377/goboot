package logger // call loggers

import (
	"logger/errorlogger"
	"logger/infologger"
)

func main() {
	// error log
	errorlogger.LogError("This is an error message!")

	// info log
	infologger.LogInfo("This is an info message!")
}
