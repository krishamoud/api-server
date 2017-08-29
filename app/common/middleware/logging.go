// Package middleware is reponsible for all api middlewares
package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// LoggingHandler prints the method, url, status code, and latency of each
// request
func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lrw := newLoggingResponseWriter(w)
		t1 := time.Now()
		next.ServeHTTP(lrw, r)
		t2 := time.Now()
		statusCode := lrw.statusCode
		l := log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL.String(),
			"statusCode": statusCode,
			"latency":    t2.Sub(t1),
		})
		switch {
		case statusCode >= 200 && statusCode < 300:
			l.Info("Successful API Request")
		case statusCode >= 300 && statusCode < 400:
			l.Warn("Redirect API Request")
		case statusCode >= 400:
			l.Error("Unsuccessful API Request")
		}
	}
	return http.HandlerFunc(fn)
}
