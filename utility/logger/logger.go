package logger

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"sns-barko/config"
	"sns-barko/constant"
	"time"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(cfg *config.Config) {
	logFile := initLogFile()
	encoderConfig := initEncoderConfig()
	var core zapcore.Core
	var lvl zapcore.Level

	if cfg.Log.Env == constant.ENV_LOCAL {
		lvl = zap.InfoLevel
		coreLog := zapcore.NewCore(
			zapcore.NewJSONEncoder(*encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logFile)),
			lvl,
		)
		core = zapcore.NewTee(coreLog)
	}

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

}

func initLogFile() *os.File {
	absPath, err := filepath.Abs("./log")
	if err != nil {
		log.Fatal("Error reading given path:", err)
	}
	logFile, err := os.OpenFile(absPath+"/logFile.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	return logFile
}

func initEncoderConfig() *zapcore.EncoderConfig {
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format("02 Jan 2006 15:04:05+07:00"))
	}

	return &zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder, // Using custom time format
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func Info(ctx context.Context, msg string, field ...zapcore.Field) {
	addTraceFromCtx(ctx, &field)
	logger.Info(msg, field...)
}

func Warn(ctx context.Context, msg string, field ...zapcore.Field) {
	addTraceFromCtx(ctx, &field)
	logger.Warn(msg, field...)
}

func Error(ctx context.Context, err error, field ...zapcore.Field) {
	addTraceFromCtx(ctx, &field)
	logger.Error(err.Error(), field...)
}

func Fatal(ctx context.Context, err error, field ...zapcore.Field) {
	addTraceFromCtx(ctx, &field)
	logger.Fatal(err.Error(), field...)
}

func Sync() {
	logger.Sync()
}

type traceInfo struct {
	traceId string
	spanId  string
}

func getTraceFromCtx(ctx context.Context) (isSpanContextValid bool, t traceInfo) {
	spanCtx := trace.SpanFromContext(ctx).SpanContext()

	if spanCtx.IsValid() {
		t.traceId = spanCtx.TraceID().String()
		t.spanId = spanCtx.SpanID().String()
	}

	return spanCtx.IsValid(), t
}

func addTraceFromCtx(ctx context.Context, field *[]zapcore.Field) {
	isSpanContextValid, t := getTraceFromCtx(ctx)
	if isSpanContextValid {
		*field = append(*field, zap.Any(constant.TRACE_ID_KEY, t.traceId), zap.Any(constant.SPAN_ID_KEY, t.spanId))
	}
}
