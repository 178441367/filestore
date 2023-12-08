package file

import (
	"context"
	"filestorage/oss/internal/logic/file/driver"
	"filestorage/oss/internal/svc"
	"filestorage/oss/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

// Upload 文件上传
func (l *UploadLogic) Upload() (resp *types.UploadResp, err error) {
	file, handler, err := l.r.FormFile("file")
	defer file.Close()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error retrieving the file,异常:%s", err.Error())
		return nil, err
	}
	oss := driver.NewOss(l.svcCtx.Config.Upload.Driver, l.svcCtx, l.r)
	url, fileName, err := oss.UploadFile(handler)
	return &types.UploadResp{
		Url:      url,
		FileName: fileName,
	}, err
}
