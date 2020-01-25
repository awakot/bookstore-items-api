package logger

import "go.uber.org/zap"

import "go.uber.org/zap/zapcore"

import "fmt"

var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.FullCallerEncoder,
		},
	}

	var err error
	if log, err = logConfig.Build(); err != nil {
		fmt.Println(fmt.Errorf("error while building logger: %v", err))
	}
}

//GetLogger returns log
func GetLogger() *zap.Logger {
	return log
}

// Info prints info level logger
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

// Error prints error level logger
func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	log.Error(msg, tags...)
	log.Sync()
}
