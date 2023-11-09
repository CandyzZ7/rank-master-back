package upload_file

import (
	"os"
	"path/filepath"
	"rank-master-back/internal/config"
	"time"

	"github.com/pkg/errors"
)

const FileFormat = "20060102/150405_" // 日期/时间_文件名.扩展名

func GetFileAndKey(config config.Config, path string) (*os.File, string, error) {
	// 获取文件
	file, err := os.Open(path)
	if err != nil {
		return nil, "", errors.WithMessage(err, "打开头像文件失败")
	}

	// 获取文件名
	avatarName := filepath.Base(path)
	// 上传地址 (后端定义，日期/时间_文件名.扩展名)
	return file, config.UploadFile.Path + time.Now().Format(FileFormat) + avatarName, nil
}
