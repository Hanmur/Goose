package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level int8

//错误等级
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

//LevelStr 错误等级映射
func (level Level) LevelStr() string {
	switch level {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

//Logger 日志器结构定义
type Logger struct {
	SugarLogger *zap.SugaredLogger
}

//NewLogger 创建日志器
func NewLogger(logPath string) (*Logger, error) {
	writeSyncer := getLogWriter(logPath)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	//基于设置后的core创建日志器
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	sugarLogger := logger.Sugar()

	return &Logger{SugarLogger: sugarLogger}, nil
}

//getEncoder 获取编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//getLogWriter 获取io流写入器
func getLogWriter(filepath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    600,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

//log 统一格式输出log
func (logger *Logger) log(level Level, message string) {
	switch level {
	case LevelDebug:
		logger.SugarLogger.Debugf(message)
	case LevelInfo:
		logger.SugarLogger.Infof(message)
	case LevelWarn:
		logger.SugarLogger.Warnf(message)
	case LevelError:
		logger.SugarLogger.Errorf(message)
	case LevelFatal:
		logger.SugarLogger.Fatalf(message)
	case LevelPanic:
		logger.SugarLogger.Panicf(message)
	}
}

//logF 统一格式输出logf
func (logger *Logger) logF(level Level, template string, args []interface{}) {
	switch level {
	case LevelDebug:
		logger.SugarLogger.Debugf(template, args)
	case LevelInfo:
		logger.SugarLogger.Infof(template, args)
	case LevelWarn:
		logger.SugarLogger.Warnf(template, args)
	case LevelError:
		logger.SugarLogger.Errorf(template, args)
	case LevelFatal:
		logger.SugarLogger.Fatalf(template, args)
	case LevelPanic:
		logger.SugarLogger.Panicf(template, args)
	}
}


//以下是对具体日志方法的实现

func (logger *Logger) Debug(message string){
	logger.log(LevelDebug, message)
}

func (logger *Logger) DebugF(template string, args ...interface{}) {
	logger.log(LevelDebug, fmt.Sprintf(template, args...))
}

func (logger *Logger) Info(message string){
	logger.log(LevelInfo, message)
}

func (logger *Logger) InfoF(template string, args ...interface{}) {
	logger.log(LevelInfo, fmt.Sprintf(template, args...))
}

func (logger *Logger) Warn(message string){
	logger.log(LevelWarn, message)
}

func (logger *Logger) WarnF(template string, args ...interface{}) {
	logger.log(LevelWarn, fmt.Sprintf(template, args...))
}

func (logger *Logger) Error(message string){
	logger.log(LevelError, message)
}

func (logger *Logger) ErrorF(template string, args ...interface{}) {
	logger.log(LevelError, fmt.Sprintf(template, args...))
}

func (logger *Logger) Fatal(message string){
	logger.log(LevelFatal, message)
}

func (logger *Logger) FatalF(template string, args ...interface{}) {
	logger.log(LevelFatal, fmt.Sprintf(template, args...))
}

func (logger *Logger) Panic(message string){
	logger.log(LevelPanic, message)
}

func (logger *Logger) PanicF(template string, args ...interface{}) {
	logger.log(LevelPanic, fmt.Sprintf(template, args...))
}
