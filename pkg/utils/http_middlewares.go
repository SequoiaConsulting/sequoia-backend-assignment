package utils

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ashutoshgngwr/sequoia-backend-assignment/pkg/service"
	"github.com/gorilla/mux"
)

// HTTPContentTypeMiddleware checks if the client accepts 'application/json' and
// sets the 'Content-type' header for all the requests
func HTTPContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accepts := r.Header.Get("Accept")
		if accepts != "*/*" && accepts != "application/json" {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// HTTPAuthMiddleware checks the session token presented with the request and adds the claims
// to its context.
func HTTPAuthMiddleware(svc service.UserService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/users/login" || r.URL.Path == "/users/signup" {
				next.ServeHTTP(w, r)
				return
			}

			claims, err := svc.Authenticate(r.Header.Get("Authorization"))
			if err == nil {
				ctx := context.WithValue(r.Context(), service.SessionClaims{}, claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				resp, _ := json.Marshal(map[string]string{"error": "login required"})
				w.WriteHeader(http.StatusForbidden)
				w.Write(resp)
			}
		})
	}
}
