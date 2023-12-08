package file

import (
	"filestorage/oss/internal/logic/file"
	"filestorage/oss/internal/svc"
	"filestorage/oss/response"
	"net/http"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := file.NewUploadLogic(r.Context(), svcCtx, r)
		resp, err := l.Upload()
		response.Response(w, r, resp, err)
	}
}
