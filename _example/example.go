package main

import "github.com/aellwein/slf4go"
import _ "github.com/aellwein/slf4go-zap-adaptor"

func main() {
	// do this only once
	slf4go.GetLoggerFactory().SetLoggingParameters(slf4go.LoggingParameters{
		"development": true,
	})

	// this is done for every logger instance
	logger := slf4go.GetLogger("example")
	logger.Info("this is just an example.")
	logger.Warn("Don't take it too serious.")
}
