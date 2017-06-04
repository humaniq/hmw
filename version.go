package hmw

import "net/http"

// Version is a middleware function that appends the cashew
// version information to the HTTP response. This is intended
// for debugging and troubleshooting or something. - manbearpig (kgan)
func Version(version string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("version", "v"+version)
			h.ServeHTTP(w, r)
		})
	}
}
