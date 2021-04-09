package middleware

import (
	"gempages-admin-server/src/pkg/helpers"
	"net/http"
	"strings"
)

// CalculateExcute ...
func CalculateExcute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api/") {
			defer helpers.Elapsed(r)()
		}

		next.ServeHTTP(w, r)
		return
	})
}
