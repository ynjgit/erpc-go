package file

import (
	"fmt"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/ynjgit/erpc-go/log"
)

func init() {
	log.RegisterWriter("file", &fileWriter{})
}

type fileWriteConfig struct {
	LogPath    string `yaml:"log_path"`
	FileName   string `yaml:"filename"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
	MaxSize    int    `yaml:"max_size"`
	Compress   bool   `yaml:"compress"`
}

type fileWriter struct {
}

func (w *fileWriter) Setup(wc log.WriterConfig) (zapcore.Core, error) {
	fwConfig := &fileWriteConfig{}
	err := wc.WriteConfig.Decode(fwConfig)
	if err != nil {
		return nil, fmt.Errorf("filewriter setup err:%s", err)
	}

	fixFileWriteConfig(fwConfig)

	level, ok := log.LevelMap[wc.Level]
	if !ok {
		level = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if wc.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	logPath := filepath.Join(fwConfig.LogPath, fwConfig.FileName)
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    fwConfig.MaxSize,
		MaxAge:     fwConfig.MaxAge,
		MaxBackups: fwConfig.MaxBackups,
		LocalTime:  true,
		Compress:   fwConfig.Compress,
	})

	core := zapcore.NewCore(
		encoder,
		writer,
		level,
	)

	return core, nil
}

func fixFileWriteConfig(fwc *fileWriteConfig) {
	if fwc.LogPath == "" {
		fwc.LogPath = "./logs"
	}

	if fwc.FileName == "" {
		fwc.FileName = "server.log"
	}

	if fwc.MaxSize <= 0 {
		fwc.MaxSize = 10 // default 10mb
	}

	if fwc.MaxAge <= 0 {
		fwc.MaxAge = 7 // default 7 days
	}

	if fwc.MaxBackups <= 0 {
		fwc.MaxBackups = 10 // default 10 backups
	}
}
