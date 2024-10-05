package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerZp *zap.SugaredLogger
)

const (
	DebugLogLevel = "DEBUG"
	InfoLogLevel  = "INFO"
	WarnLogLevel  = "WARN"
	ErrorLogLevel = "ERROR"
	FatalLogLevel = "FATAL"
)

var logLevels = map[string]zapcore.Level{
	DebugLogLevel: zap.DebugLevel,
	InfoLogLevel:  zap.InfoLevel,
	WarnLogLevel:  zap.WarnLevel,
	ErrorLogLevel: zap.ErrorLevel,
	FatalLogLevel: zap.FatalLevel,
}

const (
	defaultLevel = zap.InfoLevel
)

func init() {
	NewLogger("")
}

func NewLogger(level string, opts ...zap.Option) {
	if loggerZp != nil {
		return
	}

	logLevel := defaultLevel
	if level, ok := logLevels[level]; ok {
		logLevel = level
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)

	loggerZp = zap.New(core, opts...).Sugar()
}

func Info(args ...interface{}) {
	loggerZp.Info(args)
}

func Infof(template string, args ...interface{}) {
	loggerZp.Infof(template, args...)
}

func Debug(args ...interface{}) {
	loggerZp.Debug(args)
}

func Warn(args ...interface{}) {
	loggerZp.Warn(args)
}

func Error(args ...interface{}) {
	loggerZp.Error(args)
}

func Errorf(template string, args ...interface{}) {
	loggerZp.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	loggerZp.Fatal(args)
}

func Fatalf(template string, args ...interface{}) {
	loggerZp.Fatalf(template, args)
}

func Logger() *zap.SugaredLogger {
	return loggerZp
}
