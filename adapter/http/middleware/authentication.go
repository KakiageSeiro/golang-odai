package middleware

import (
	"golang-odai/adapter/http/session"
	"net/http"
)

// AuthenticationMiddleware checks login sesison
func AuthenticationMiddleware(sess *session.Session) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			_, err := sess.GetUser(r)
			if err != nil {
				panic(err)
				http.Redirect(w, r, "/signin", http.StatusFound)
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
