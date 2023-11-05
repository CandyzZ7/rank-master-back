package oss

import (
	"io"
	"rank-master-back/internal/svc"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
)

func UploadFile(svc *svc.ServiceContext, objectKey string, file io.Reader) error {
	bucket, err := GetBucket(svc)
	if err != nil {
		return errors.Wrap(err, "get bucket error")
	}
	err = PutObject(bucket, objectKey, file)
	if err != nil {
		return errors.Wrap(err, "put object error")
	}
	return nil
}

func GetBucket(svc *svc.ServiceContext) (*oss.Bucket, error) {
	bucket, err := svc.Oss.Bucket(svc.Config.UploadFile.AliYunOss.BucketName)
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
