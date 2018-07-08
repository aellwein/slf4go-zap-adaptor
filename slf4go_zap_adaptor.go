package slf4go_zap_adaptor

import (
	"errors"
	"fmt"
	"github.com/aellwein/slf4go"
	"go.uber.org/zap"
)

type zapLoggerFactory struct {
	logger *zap.Logger
	level  slf4go.LogLevel
}

// Creates by default a production sugared logger
func (f *zapLoggerFactory) GetLogger(name string) slf4go.Logger {
	logger := &zapLogger{}
	logger.sugar = f.logger.Sugar()
	logger.LoggerAdaptor.SetLevel(f.level)
	return logger
}

func (f *zapLoggerFactory) SetDefaultLogLevel(lvl slf4go.LogLevel) {
	f.level = lvl
}

func (f *zapLoggerFactory) GetDefaultLogLevel() slf4go.LogLevel {
	return f.level
}

type zapLogger struct {
	slf4go.LoggerAdaptor
	sugar *zap.SugaredLogger
}

func (l *zapLogger) Trace(args ...interface{}) {
	// trace maps to debug
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelTrace {
		l.sugar.Debug(args...)
	}
}

func (l *zapLogger) Tracef(format string, args ...interface{}) {
	// trace maps to debug
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelTrace {
		l.sugar.Debugf(format, args...)
	}
}

func (l *zapLogger) Debug(args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelDebug {
		l.sugar.Debug(args...)
	}
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelDebug {
		l.sugar.Debugf(format, args...)
	}
}

func (l *zapLogger) Info(args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelInfo {
		l.sugar.Info(args...)
	}
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelInfo {
		l.sugar.Infof(format, args...)
	}
}

func (l *zapLogger) Warn(args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelWarn {
		l.sugar.Warn(args...)
	}
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelWarn {
		l.sugar.Warnf(format, args...)
	}
}

func (l *zapLogger) Error(args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelError {
		l.sugar.Error(args...)
	}
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelError {
		l.sugar.Errorf(format, args...)
	}
}

func (l *zapLogger) Fatal(args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelFatal {
		l.sugar.Fatal(args...)
	}
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelFatal {
		l.sugar.Fatalf(format, args...)
	}
}

func (l *zapLogger) Panic(args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelPanic {
		l.sugar.Panic(args...)
	}
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	if l.LoggerAdaptor.GetLevel() <= slf4go.LevelPanic {
		l.sugar.Panicf(format, args...)
	}
}

func newZapLoggerFactory() slf4go.LoggerFactory {
	factory := &zapLoggerFactory{level: slf4go.LevelInfo}
	factory.logger, _ = zap.NewProduction()
	return factory
}

func (f *zapLoggerFactory) SetLoggingParameters(params slf4go.LoggingParameters) error {
	for k, v := range params {
		switch k {
		case "development":
			if dev, ok := v.(bool); !ok {
				return errors.New("invalid type for parameter 'development', should be of type bool")
			} else {
				if dev {
					f.logger, _ = zap.NewDevelopment()
				}
			}

		case "options":
			if options, ok := v.([]zap.Option); !ok {
				return errors.New("invalid type for parameter 'options', should be of type []zap.Option")
			} else {
				f.logger = f.logger.WithOptions(options...)
			}

		case "fields":
			if fields, ok := v.([]zap.Field); !ok {
				return errors.New("invalid type for parameter 'fields', should be of type []zap.Field")
			} else {
				f.logger = f.logger.With(fields...)
			}
		default:
			return errors.New(fmt.Sprintf("unsupported parameter: %v", k))
		}
	}
	return nil
}
