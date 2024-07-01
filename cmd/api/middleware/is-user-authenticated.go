package middleware

import (
	"fmt"
	"net/http"

	"github.com/joaocansi/essay-api/cmd/api/httpres"
)

func IsUserAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			httpres.New(w).Error(http.StatusUnauthorized, "unauthorized account")
			return
		}

		fmt.Print(token)
		next.ServeHTTP(w, r)
	})
}
