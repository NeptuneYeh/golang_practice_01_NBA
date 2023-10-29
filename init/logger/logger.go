package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Any         = zap.Any
)

type LoggerModule struct {
	zapLogger *zap.Logger
}

var SystemLogger *LoggerModule

func NewModule(initialFields map[string]interface{}) *LoggerModule {
	loggerModule := &LoggerModule{
		zapLogger: InitLoggerConfig(initialFields),
	}
	SystemLogger = loggerModule
	return loggerModule
}

func InitLoggerConfig(initialFields map[string]interface{}) *zap.Logger {
	zapConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(getLogLevelFromViper()),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:   "level",
			TimeKey:    "timestamp",
			MessageKey: "message",
			// FunctionKey: "function",
			// StacktraceKey:  "stacktrace", // will log trace when level is under warning
			CallerKey:      "file",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		},
		// OutputPaths:      []string{"stdout", "../logs/log.log"},
		OutputPaths:      []string{"stdout", "../logs/log.log"},
		ErrorOutputPaths: []string{"stdout", "../logs/log.log"},
		InitialFields: map[string]interface{}{
			"process_name": "Golang NBA Data Scraper",
			"version":      "1.0.0",
		},
	}
	if len(initialFields) > 0 {
		for key, value := range initialFields {
			zapConfig.InitialFields[key] = value
		}
	}
	newLogger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	return newLogger
}

func getLogLevelFromViper() zapcore.Level {
	switch strings.ToLower("DEBUG") {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func (l *LoggerModule) Run(osChannel chan os.Signal) error {
	return nil
}

func (l *LoggerModule) Shutdown() error {
	l.Info("Shutting down LoggerModule...")
	l.zapLogger.Sync()
	return nil
}

func (l *LoggerModule) Debug(message string, fields ...zap.Field) {
	l.zapLogger.WithOptions(zap.AddCallerSkip(1)).Debug(message, fields...)
}

func (l *LoggerModule) DebugNoSkip(message string, fields ...zap.Field) {
	l.zapLogger.Debug(message, fields...)
}

func (l *LoggerModule) Info(message string, fields ...zap.Field) {
	l.zapLogger.WithOptions(zap.AddCallerSkip(1)).Info(message, fields...)
}

func (l *LoggerModule) Warn(message string, fields ...zap.Field) {
	l.zapLogger.WithOptions(zap.AddCallerSkip(1)).Warn(message, fields...)
}

func (l *LoggerModule) Error(message string, fields ...zap.Field) {
	l.zapLogger.WithOptions(zap.AddCallerSkip(1)).Error(message, fields...)
}

func (l *LoggerModule) ErrorWithDetail(message string, errorCode string, errorMessage string, errorTarget string) {
	l.zapLogger.WithOptions(zap.AddCallerSkip(1)).Error(message,
		String("error_code", errorCode),
		String("error_message", errorMessage),
		String("error_target", errorTarget))
}

func (l *LoggerModule) ErrorWithStackTrace(message string, fields ...zap.Field) {
	fields = append(fields, zap.Stack("stack_trace"))
	l.zapLogger.WithOptions(zap.AddCallerSkip(1)).Error(message, fields...)
}
