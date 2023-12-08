package middleware

import "net/http"

type CorsMiddleware struct {
}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		if r.Method == "OPTIONS" {
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Passthrough to next handler if need
		next(w, r)
	}
}
