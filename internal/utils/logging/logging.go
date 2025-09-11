// Package logging provides logging functionality for the application.
package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func InitLogger(cfg *Config) {
	var cores []zapcore.Core

	// Console logging
	if cfg.Development {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleWriter := zapcore.AddSync(os.Stdout)
		consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, getZapLevel(cfg.Level))
		cores = append(cores, consoleCore)
	} else {
		consoleEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		consoleWriter := zapcore.Lock(zapcore.AddSync(os.Stdout))
		consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, getZapLevel(cfg.Level))
		cores = append(cores, consoleCore)
	}

	// File logging
	if cfg.FilePath != "" {
		fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		fileWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.FilePath,
			MaxSize:    100, // megabytes
			MaxAge:     7,   // days
			MaxBackups: 3,
			Compress:   true,
		})
		fileLevel := getZapLevel(cfg.FileLevel)
		if cfg.FileLevel == "" {
			fileLevel = getZapLevel(cfg.Level)
		}
		fileCore := zapcore.NewCore(fileEncoder, fileWriter, fileLevel)
		cores = append(cores, fileCore)
	}

	core := zapcore.NewTee(cores...)
	logger = zap.New(core, zap.AddCaller())
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func GetLogger() *zap.Logger {
	return logger
}