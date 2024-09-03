package middlewares

import (
	"gitlab.com/roneeSoft/integrator/pkg/shared/utils"
	"net/http"
	"time"
)

func LoggingMiddleware(logger *utils.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logger.Info("Started " + r.Method + " " + r.URL.Path)

			next.ServeHTTP(w, r)

			logger.Info("Completed in " + time.Since(start).String())
		})
	}
}
