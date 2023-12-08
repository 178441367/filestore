package driver

import (
	"filestorage/oss/internal/svc"
	"mime/multipart"
	"net/http"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss(ossType string, svcCtx *svc.ServiceContext, req *http.Request) OSS {
	switch ossType {
	case "local":
		return &LocalDriver{
			r:      req,
			svcCtx: svcCtx,
		}
	default:
		return &LocalDriver{
			r:      req,
			svcCtx: svcCtx,
		}
	}
}
