package xoss

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"

	"rank-master-back/internal/config"
)

func NewOssClient(c config.Config) (*oss.Client, error) {
	ossClient, err := oss.New(c.UploadFile.AliYunOss.Endpoint, c.UploadFile.AliYunOss.AccessKeyId, c.UploadFile.AliYunOss.AccessKeySecret,
		oss.Timeout(c.UploadFile.AliYunOss.ConnectTimeout, c.UploadFile.AliYunOss.ReadWriteTimeout))
	if err != nil {
		return nil, err
	}
	return ossClient, nil
}

func UploadFile(client *oss.Client, bucketName string, objectKey string, file io.Reader) error {
	bucket, err := GetBucket(client, bucketName)
	if err != nil {
		return errors.Wrap(err, "get bucket error")
	}
	err = PutObject(bucket, objectKey, file)
	if err != nil {
		return errors.Wrap(err, "put object error")
	}
	return nil
}

func GetBucket(client *oss.Client, bucketName string) (*oss.Bucket, error) {
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, errors.WithMessage(err, "upload_file get bucket error")
	}
	return bucket, nil
}

func PutObject(bucket *oss.Bucket, objectKey string, file io.Reader) error {
	err := bucket.PutObject(objectKey, file)
	if err != nil {
		return errors.WithMessage(err, "upload_file put object error")
	}
	return nil
}
