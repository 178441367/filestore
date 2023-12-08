package driver

import (
	"errors"
	"filestorage/oss/internal/svc"
	"filestorage/oss/utils"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"mime/multipart"
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

func (l LocalDriver) UploadFile(fileheader *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(fileheader.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(fileheader.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	date := time.Now().Format("200601")
	dir := fmt.Sprintf("%s/%s", l.svcCtx.Config.Upload.Dir, date)
	localPath := filepath.Join(dir, filename)
	var scheme string
	if l.r.URL.Scheme == "" {
		scheme = "http://"
	} else {
		scheme = l.r.URL.Scheme + "://"
	}

	host := scheme + l.r.Host
	url := fmt.Sprintf("%s%s/%s/%s", host, l.svcCtx.Config.Upload.Prefix, date, filename)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logx.Errorf("Failed to create directory,异常:%s", err.Error())
		return "", "", err
	}
	// 创建一个新的文件来保存上传的文件
	uploadedFile, err := os.Create(localPath)
	if err != nil {
		logx.Errorf("Error creating the filee,异常:%s", err.Error())
		return "", "", err
	}
	defer uploadedFile.Close()
	file, openError := fileheader.Open() // 读取文件
	if openError != nil {
		logx.Errorf("function file.Open() failed", err.Error())
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	// 将上传的文件拷贝到新文件中
	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		logx.Errorf("Error copying the file,异常:%s", err.Error())
		return "", "", err
	}

	logx.Infof("File uploaded successfully: %s", filename)
	return url, filename, err
}
