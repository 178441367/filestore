package middleware

import (
	"errors"
	"filestorage/oss/global"
	"filestorage/oss/response"
	"fmt"
	"net/http"
)

type TokenMiddleware struct {
}

func NewTokenMiddleware() *TokenMiddleware {
	return &TokenMiddleware{}
}

func (m *TokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		fmt.Println("TokenMiddleware")
		token := r.Header.Get("Authorization")
		if token != global.Conf.Authorization {
			response.Response(w, r, nil, errors.New("token is error"))
			return
		}
		next(w, r)
	}
}
