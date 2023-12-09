package driver

import (
	"filestorage/oss/global"
	"filestorage/oss/internal/svc"
	"filestorage/oss/internal/types"
	"net/http"
)

type OSS interface {
	UploadFile(*types.UploadReq) (string, string, error)
}

// NewOss OSS的实例化方法
func NewOss(svcCtx *svc.ServiceContext, req *http.Request) OSS {
	switch global.Conf.Upload.Driver {
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
