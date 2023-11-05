package local

import (
	"io"
	"os"
	"path/filepath"
	"rank-master-back/internal/config"
	"rank-master-back/internal/pkg/upload_file"

	"github.com/pkg/errors"
)

func Upload(config config.Config, path string) (string, error) {
	// 获取头像文件和上传地址
	file, key, err := upload_file.GetFileAndKey(config, path)
	if err != nil {
		return "", errors.WithMessage(err, "获取头像文件和上传地址失败")
	}
	// 上传头像
	dir := filepath.Dir(key)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", errors.WithMessage(err, "创建文件夹失败")
	}
	out, err := os.Create(key)
	if err != nil {
		return "", errors.WithMessage(err, "创建文件失败")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(out)
	_, err = io.Copy(out, file)
	if err != nil {
		return "", errors.WithMessage(err, "上传头像失败")
	}
	return key, nil
}
