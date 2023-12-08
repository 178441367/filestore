package svc

import (
	"filestorage/oss/internal/config"
	"filestorage/oss/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	TokenMiddleware rest.Middleware
	CorsMiddleware  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		TokenMiddleware: middleware.NewTokenMiddleware().Handle,
		CorsMiddleware:  middleware.NewCorsMiddleware().Handle,
	}
}
