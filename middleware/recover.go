package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logrus.WithFields(logrus.Fields{
					"error":      err,
					"stacktrace": string(debug.Stack()),
					"path":       r.URL.Path,
					"method":     r.Method,
				}).Error("Recovered from panic")

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
