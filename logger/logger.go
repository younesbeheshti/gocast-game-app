package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
)

var Logger *zap.Logger

var once sync.Once

func init() {
	once.Do(func() {
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		defaultEncoder := zapcore.NewJSONEncoder(config)
		//logFile, _ := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		//writer := zapcore.AddSync(logFile)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:  "./logs/log.json",
			LocalTime: false,
			MaxSize:   10, // megabytes
			//MaxBackups: 2,
			MaxAge:   20,   //days
			Compress: true, // disabled by default
		})
		stdOutWriter := zapcore.AddSync(os.Stdout)
		defaultLogLevel := zapcore.DebugLevel
		infoLogLevel := zapcore.InfoLevel
		core := zapcore.NewTee(
			zapcore.NewCore(defaultEncoder, writer, defaultLogLevel),
			zapcore.NewCore(defaultEncoder, stdOutWriter, infoLogLevel),
		)
		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	})

}
