// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"filestorage/oss/internal/handler/file"
	"filestorage/oss/internal/svc"
	"net/http"


	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: file.UploadHandler(serverCtx),
			},
		},
	)
}
