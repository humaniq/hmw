package hmw

import (
	"fmt"
	"net/http"
	"strings"
)

// isContentType validates the Content-Type header matches the supplied
// contentType. That is, its type and subtype match.
func isContentType(h http.Header, contentType string) bool {
	ct := h.Get("Content-Type")
	if i := strings.IndexRune(ct, ';'); i != -1 {
		ct = ct[0:i]
	}
	return ct == contentType
}

// ContentType wraps and returns a http.Handler, validating the request
// content type is compatible with the contentTypes list. It writes a HTTP 415
// error if that fails.
// Only PUT, POST, and PATCH requests are considered.
func ContentType(contentTypes ...string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !(r.Method == "PUT" || r.Method == "POST" || r.Method == "PATCH") {
				h.ServeHTTP(w, r)
				return
			}

			for _, ct := range contentTypes {
				if isContentType(r.Header, ct) {
					h.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, fmt.Sprintf("Unsupported content type %q; expected one of %q", r.Header.Get("Content-Type"), contentTypes), http.StatusUnsupportedMediaType)
		})
	}
}
