package file

import (
	"api/oss/utils"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"api/oss/internal/svc"
	"api/oss/internal/types"

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
// 参考：https://github.com/zeromicro/zero-examples/blob/main/monolithic/internal/logic/uploadlogic.go
func (l *UploadLogic) Upload() (resp *types.UploadResp, err error) {

	file, handler, err := l.r.FormFile("file")
	defer file.Close()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error retrieving the file,异常:%s", err.Error())
		return nil, err
	}
	// 读取文件后缀
	ext := path.Ext(handler.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(handler.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	date := time.Now().Format("200601")
	dir := fmt.Sprintf("%s/%s", l.svcCtx.Config.Upload.Dir, date)
	localPath := filepath.Join(dir, filename)
	host := l.r.Host
	url := fmt.Sprintf("%s%s/%s/%s", host, l.svcCtx.Config.Upload.Prefix, date, filename)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Failed to create directory,异常:%s", err.Error())
		return nil, err
	}
	// 创建一个新的文件来保存上传的文件
	uploadedFile, err := os.Create(localPath)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error creating the filee,异常:%s", err.Error())
		return nil, err
	}
	defer uploadedFile.Close()

	// 将上传的文件拷贝到新文件中
	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error copying the file,异常:%s", err.Error())
		return nil, err
	}

	logx.WithContext(l.ctx).Infof("File uploaded successfully: %s", filename)
	return &types.UploadResp{
		Url: url,
	}, nil
}
