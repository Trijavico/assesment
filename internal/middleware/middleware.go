package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Trijavico/ensolvers-assesment/internal/model/dto"
	"github.com/Trijavico/ensolvers-assesment/internal/service"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(middlewares ...Middleware) Middleware {

	return func(next http.Handler) http.Handler {
		for _, middleware := range middlewares {
			next = middleware(next)
		}

		return next
	}
}

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtService := service.JwtService{}
		value := r.Header.Get("Authorization")

		if !strings.Contains(value, "bearer ") {
			http.Error(w, service.UnAuthorized.Error(), http.StatusUnauthorized)
			return
		}
		if len(value) <= len("bearer ") {
			http.Error(w, service.UnAuthorized.Error(), http.StatusUnauthorized)
			return
		}

		token := value[len("bearer "):]

		claims, err := jwtService.ParseToken(token)
		if err != nil {
			http.Error(w, service.UnAuthorized.Error(), http.StatusUnauthorized)
			return
		}

		userID := claims["id"].(float64)
		if userID < 0 {
			http.Error(w, service.UnAuthorized.Error(), http.StatusUnauthorized)
			return
		}

		authUser := dto.AuthUser{
			ID: uint(userID),
		}

		ctxWithUser := context.WithValue(r.Context(), "user", authUser)
		r = r.WithContext(ctxWithUser)

		next.ServeHTTP(w, r)
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
