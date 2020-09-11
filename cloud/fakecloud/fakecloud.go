/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/11 11:35
* @Description: 一个虚假的云存储，其实只是简单地等待一段时间在内存中标记文件已上传
***********************************************************************/

package fakecloud

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"
)

// FakeCloud
type FakeCloud struct {
	bucket string
	uploaded map[string]bool	// 文件已上传
}

// Upload 以二进制流格式（默认的ContentType）上传文件
func (fc *FakeCloud) Upload(ctx context.Context, path string) error {
	_, filename := filepath.Split(path)

	// 模拟上传时模拟1MB/s的上传速度
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	size := info.Size() * 10	// 测试时文件大小模拟成10倍实际大小
	mb := size / (1024 * 1024)

	time.Sleep(time.Duration(mb) * time.Second)
	fc.uploaded[filename] = true
	log.Printf("Successfully uploaded %s\n", path)
	return nil
}

// NewFakeCloud
func NewFakeCloud(ctx context.Context, cloudpath string) (*FakeCloud, error) {
	m := new(FakeCloud)
	m.bucket = cloudpath
	m.uploaded = make(map[string]bool)

	log.Printf("Successfully created %s\n", m.bucket)

	return m, nil
}

