package log

import (
	"math/rand"
	"net/http"
	"slices"
)

func getTraceIdOrCreate(r *http.Request) string {
	if traceId := r.URL.Query().Get("trace_id"); len(traceId) > 0 {
		return traceId
	}
	if traceId := r.Header.Get("x-trace-id"); len(traceId) > 0 {
		return traceId
	}
	traceId, _ := NewUuid().Generate()
	return traceId
}

var LEVELS = []string{"verbose", "debug", "info", "warn", "error", "fatal", "silent"}

func getLevel(r *http.Request) *string {
	level := getLevelString(r)
	if slices.Contains(LEVELS, level) {
		return &level
	}
	return nil
}

func getLevelString(r *http.Request) string {
	if level := r.URL.Query().Get("log_level"); len(level) > 0 {
		return level
	}
	if level := r.Header.Get("x-log-level"); len(level) > 0 {
		return level
	}
	return ""
}

func percentFloat(percent float64) bool {
	if percent < 0 || percent > 100 {
		panic("Percent must by in [0.0, 100.0]")
	}
	value := rand.Float64() * 100
	return value <= percent
}
