package log

import (
	"context"
	"fmt"
	"net/http"
	"runtime"

	loggergo "github.com/nextmillenniummedia/logger-go"
)

type key string

var LOGGER_KEY key = "nm_logger"
var TRACE_ID key = "nm_trace_id"
var LEVEL key = "nm_level"

var (
	_logger   loggergo.ILogger
	_level    string
	_sampling float64
)

func New(config LoggerConfig) loggergo.ILogger {
	_logger = loggergo.New().LevelHuman(config.Level)
	_sampling = config.Sampling
	_level = config.Level
	if config.Pretty {
		_logger.Pretty()
	}
	_logger.Clone().From("logger").Info("Init")
	return _logger
}

func SetToRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceId := getTraceIdOrCreate(r)
		logger := cloneWithSampling().Params("trace_id", traceId)
		level := getLevel(r)
		if level != nil {
			logger.LevelHuman(*level)
		}
		if level == nil {
			level = &_level
		}
		w.Header().Set("x-trace-id", traceId)
		ctx := setToContext(r.Context(), logger, traceId, *level)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func GetFromContext(ctx context.Context, from string) loggergo.ILogger {
	logger, _, _ := getFromContext(ctx, from)
	return logger
}

func GetFromRequest(r *http.Request, from string) loggergo.ILogger {
	return GetFromContext(r.Context(), from)
}

func GetTraceIdFromContext(ctx context.Context) string {
	_, trace_id, _ := getFromContext(ctx, "")
	return trace_id
}

func cloneWithSampling() loggergo.ILogger {
	logger := _logger.Clone()
	if percentFloat(_sampling) {
		return logger
	}
	return logger.Level(loggergo.LOG_SILENT)
}

func setToContext(ctx context.Context, logger loggergo.ILogger, traceId, level string) context.Context {
	if ctx != nil {
		ctx = context.WithValue(ctx, LOGGER_KEY, logger)
		ctx = context.WithValue(ctx, TRACE_ID, traceId)
		ctx = context.WithValue(ctx, LEVEL, level)
	}
	return ctx
}

func getFromContext(ctx context.Context, from string) (logger loggergo.ILogger, traceId string, level string) {
	loggerValue := ctx.Value(LOGGER_KEY)
	if loggerValue == nil {
		_, filePath, strNum, _ := runtime.Caller(1)
		line := fmt.Sprintf("%v:%v", filePath, strNum)
		logger = cloneWithSampling().From(from).Info("Logger not found in context", "line", line)
	} else {
		logger = loggerValue.(loggergo.ILogger).Clone().From(from)
	}

	traceIdValue := ctx.Value(TRACE_ID)
	if traceIdValue != nil {
		traceId = traceIdValue.(string)
	} else {
		traceId = ""
	}

	levelValue := ctx.Value(LEVEL)
	if levelValue != nil {
		level = levelValue.(string)
	} else {
		level = _level
	}

	return logger, traceId, level
}
