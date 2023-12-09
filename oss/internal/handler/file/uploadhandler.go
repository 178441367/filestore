package file

import (
	"filestorage/oss/internal/logic/file"
	"filestorage/oss/internal/svc"
	"filestorage/oss/internal/types"
	"filestorage/oss/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, err)
			return
		}
		l := file.NewUploadLogic(r.Context(), svcCtx, r)
		resp, err := l.Upload(&req)
		response.Response(w, r, resp, err)
	}
}
