package driver

import (
	"filestorage/oss/internal/svc"
	"filestorage/oss/internal/types"
	"filestorage/oss/utils"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type LocalDriver struct {
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func (l LocalDriver) UploadFile(req *types.UploadReq) (string, string, error) {
	// 读取文件后缀
	file, handler, err := l.r.FormFile("file")
	defer file.Close()
	var (
		fileDir  string
		url      string
		fileName string
	)
	ext := path.Ext(handler.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(handler.Filename, ext)
	// 拼接新文件名
	if req.FileNameCreate == 1 {
		name = utils.MD5V([]byte(name))
		fileName = name + "_" + time.Now().Format("20060102150405") + ext
	} else {
		fileName = name + ext
	}
	dir := req.Dir
	//date := time.Now().Format("200601")
	if dir == "" {
		dir = time.Now().Format("200601")
	}
	fileDir = fmt.Sprintf("%s/%s", l.svcCtx.Config.Upload.Dir, dir)
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		logx.Errorf("Failed to create directory,异常:%s", err.Error())
		return "", "", err
	}
	localPath := filepath.Join(fileDir, fileName)
	// 创建一个新的文件来保存上传的文件
	uploadedFile, err := os.Create(localPath)
	if err != nil {
		logx.Errorf("Error creating the filee,异常:%s", err.Error())
		return "", "", err
	}
	defer uploadedFile.Close()
	// 将上传的文件拷贝到新文件中
	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		logx.Errorf("Error copying the file,异常:%s", err.Error())
		return "", "", err
	}

	logx.Infof("File uploaded successfully: %s", fileName)

	var scheme string
	if l.r.URL.Scheme == "" {
		scheme = "http://"
	} else {
		scheme = l.r.URL.Scheme + "://"
	}
	host := scheme + l.r.Host

	url = fmt.Sprintf("%s%s/%s/%s", host, l.svcCtx.Config.Upload.Prefix, dir, fileName)
	return url, fileName, err
}
