/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 14:01
* @Description: Minio对象存储客户端的封装
***********************************************************************/

package minio

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/azd1997/dumper2cloud/conf"
	miniogo "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type BucketExistsErr struct {
	bucket string
}

func (e BucketExistsErr) Error() string {
	return fmt.Sprintf("bucket %s already exists", e.bucket)
}

// Minio minio客户端
type Minio struct {
	client *miniogo.Client
	bucket string
}

// Upload 以二进制流格式（默认的ContentType）上传文件
func (m *Minio) Upload(ctx context.Context, path string) error {
	_, filename := filepath.Split(path)
	_, err := m.client.FPutObject(ctx, m.bucket, filename, path,
		miniogo.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

// NewMinio 新建Minio客户端
func NewMinio(ctx context.Context, cloudpath string) (*Minio, error) {
	// 准备配置
	endpoint := conf.Global().Section("minio").Key("endpoint").String()
	accessKeyID := conf.Global().Section("minio").Key("access_key_id").String()
	secretAccessKey := conf.Global().Section("minio").Key("secret_access_key").String()
	useSSL, err := conf.Global().Section("minio").Key("use_ssl").Bool()
	if err != nil {
		return nil, err
	}
	bucketRegion := conf.Global().Section("minio").Key("bucket_region").String()

	// 创建minio客户端
	minioClient, err := miniogo.New(endpoint, &miniogo.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	// 创建Minio
	m := new(Minio)
	m.client = minioClient
	m.bucket = cloudpath

	// 创建Minio bucket
	err = m.client.MakeBucket(ctx, m.bucket,
		miniogo.MakeBucketOptions{Region: bucketRegion})
	if err != nil {
		// 检查bucket是否已存在
		exists, errBucketExists := minioClient.BucketExists(ctx, m.bucket)
		if errBucketExists == nil && exists {
			return m, BucketExistsErr{bucket: m.bucket}
		} else {
			return nil, err
		}
	} else {
		log.Printf("Successfully created %s\n", m.bucket)
	}

	return m, nil
}
