package middleware

import (
	"log/slog"
	"net/http"
	"sync"
	"time"
)

var (
	vlog *slog.Logger
	once sync.Once
)

func Logger(next http.Handler) http.Handler {
	once.Do(func() {
		vlog = slog.Default()
		slog.SetLogLoggerLevel(slog.LevelDebug)
	})
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := wrapWriter(w)
			next.ServeHTTP(ww, r)
			vlog.Info("http",
				"status", ww.status,
				"method", r.Method,
				"path", r.URL.EscapedPath(),
				"duration", time.Since(start),
			)
		},
	)
}

type responseWriter struct {
	http.ResponseWriter

	status      int
	wroteHeader bool
}

func wrapWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}
