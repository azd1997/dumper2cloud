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
	bucket   string
	uploaded map[string]bool // 文件已上传
}

// Upload 以二进制流格式（默认的ContentType）上传文件
func (fc *FakeCloud) Upload(ctx context.Context, path string) error {
	_, filename := filepath.Split(path)

	// 测试时，模拟上传时模拟100MB/s的上传速度，不然等控制台输出要等很久
	// 以后添加fakecloud的配置
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	size := info.Size()
	mb := size >> 20

	// 测试文件是10MB的话，等0.1s
	time.Sleep(time.Duration(mb * 10) * time.Millisecond)
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
