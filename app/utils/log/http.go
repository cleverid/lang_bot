package log

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"slices"
	"time"

	loggergo "github.com/nextmillenniummedia/logger-go"
)

var excludeLogUri = []string{"/status"}

func Http(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := GetFromRequest(r, "http")

		isExcluded := slices.Contains(excludeLogUri, r.RequestURI)
		notDebugLevel := !logger.HasLevel(loggergo.LOG_DEBUG)
		if isExcluded || notDebugLevel {
			next.ServeHTTP(w, r)
			return
		}

		startTime := time.Now()

		// For body reading
		body, bodyParsed := GetRequestBody(r)
		r.Body = body

		logger.Debug("request",
			"url", r.RequestURI,
			"method", r.Method,
			"headers", GetHeaders(r.Header),
			"payload", formatResponse(json.RawMessage(bodyParsed), logger.IsPretty()),
		)

		// For response reading
		w2 := NewCustomWriter(w)
		next.ServeHTTP(w2, r)

		logger.Debug("response",
			"time", time.Since(startTime).String(),
			"url", r.RequestURI,
			"status", w2.statusCode,
			"headers", GetHeaders(w2.ResponseWriter.Header()),
			"payload", formatResponse(w2.GetResponse(), logger.IsPretty()),
		)
	})
}

func NewCustomWriter(w http.ResponseWriter) *CustomWriter {
	return &CustomWriter{
		ResponseWriter: w,
		statusCode:     200,
		response:       make([][]byte, 0),
	}
}

type CustomWriter struct {
	http.ResponseWriter
	response   [][]byte
	statusCode int
}

func (rw *CustomWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *CustomWriter) Write(p []byte) (int, error) {
	rw.response = append(rw.response, p)
	return rw.ResponseWriter.Write(p)
}

func (rw *CustomWriter) GetResponse() json.RawMessage {
	response := string(bytes.Join(rw.response, []byte{}))
	if response == "" {
		response = "{}"
	}
	return json.RawMessage(response)
}

func GetHeaders(h http.Header) (headers map[string]any) {
	headers = make(map[string]any)
	for name, value := range h {
		headers[name] = value
	}
	return headers
}

func GetRequestBody(r *http.Request) (body io.ReadCloser, result json.RawMessage) {
	raw, _ := io.ReadAll(r.Body)
	parsed := string(raw)
	if parsed == "" {
		parsed = "{}"
	}
	r.Body.Close()
	body = io.NopCloser(bytes.NewBuffer(raw))
	return body, json.RawMessage(parsed)
}

func formatResponse(response json.RawMessage, isPretty bool) any {
	if isPretty {
		return string(response)
	}
	return response
}
