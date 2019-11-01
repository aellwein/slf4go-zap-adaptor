[![Go Report Card](https://goreportcard.com/badge/github.com/aellwein/slf4go-zap-adaptor)](https://goreportcard.com/report/github.com/aellwein/slf4go-zap-adaptor)
[![Coverage Status](https://img.shields.io/coveralls/github/aellwein/slf4go-zap-adaptor/master.svg)](https://coveralls.io/github/aellwein/slf4go-zap-adaptor?branch=master)
[![Build Status](https://img.shields.io/travis/aellwein/slf4go-zap-adaptor/master.svg)](https://travis-ci.org/aellwein/slf4go-zap-adaptor) 



# Zap adaptor for SLF4GO

This is a [zap](https://github.com/uber-go/zap) adaptor implementation for [SLF4GO](https://github.com/aellwein/slf4go).

An example usage is stupid simple:

```go

package main

import "github.com/aellwein/slf4go"
import _ "github.com/aellwein/slf4go-zap-adaptor"

func main() {
	slf4go.GetLoggerFactory().SetLoggingParameters(slf4go.LoggingParameters{
		"development": true,
	})

	logger := slf4go.GetLogger("example")
	logger.Info("this is just an example.")
	logger.Warn("Don't take it too serious.")
}

```
Note the underscore in front of the import of the SLF4GO adaptor.

You can change the logger implementation anytime, without changing the facade you are using, only by changing 
the imported adaptor.

# Logging parameters

This adaptor supports several parameters, available with ``SetLoggingParameters``:


 Parameter Key     | Value Type                        | Description
-------------------|-----------------------------------|----------------------------------
 "development"     | ``bool``                          | true, if development logging should be used 
 "options"         | ``[]zap.Option``                  | Arbitrary options to pass to the logger factory
 "fields"          | ``[]zap.Field``                   | Fields to be included 

# Development

* use ``go build ./...`` as usual.