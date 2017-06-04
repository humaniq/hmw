package hmw

import (
	"errors"
	"net/http"

	"github.com/humaniq/hmnqlog"

	"go.uber.org/zap"
)

// Recover ...
func Recover(lgr hmnqlog.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var err error
			defer func() {
				r := recover()
				if r != nil {
					switch t := r.(type) {
					case string:
						err = errors.New(t)
					case error:
						err = t
					default:
						err = errors.New("Unknown error")
					}
					lgr.Error("Captured Exception", zap.String("exception", err.Error()))
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}()
			h.ServeHTTP(w, r)
		})
	}
}
