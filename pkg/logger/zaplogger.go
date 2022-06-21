package logger

import (
	"os"
	"sail/global"
	"sail/pkg/setting"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger 初始化Logger
func InitLogger(cfg *setting.LogConf) (err error) {
	// 写入日志文件配置
	writeSyncer := getLogWriter(cfg.LogFile, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	// 日志格式
	encoder := getEncoder()
	// 日志等级
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return err
	}
	// 定义zapcore, 同时写入标准输出和日志文件
	multiSyncer := zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(encoder, multiSyncer, l)

	// 初始化
	global.Logger = zap.New(core, zap.AddCaller())
	return nil
}

// 定义日志编码格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

// 定义日志文件配置
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
