/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 16:04
* @Description: 测试cloud是否正常
***********************************************************************/

package cloud

import (
	"context"
	"fmt"
	"github.com/azd1997/dumper2cloud/cloud/minio"
	"github.com/azd1997/dumper2cloud/conf"
	"os"
	"testing"
)

// 在test/下创建临时文件
func createTestFile(name string) {
	file, err := os.Create("../test/" + name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := "abcdefghijk\n"
	for i:=0; i<100; i++ {
		file.WriteString(data)
	}
}

// 在test/下删除临时文件
func removeTestFile(name string) {
	os.Remove("../test/" + name)
}

//
func TestFakeCloud(t *testing.T) {
	filename := "test-fake-file"
	createTestFile(filename)
	defer removeTestFile(filename)

	conf.Init("../conf/config.ini")

	// 打桩
	cloudtype = func() string {
		return "fakecloud"
	}

	cld, err := NewCloud(context.TODO(), "testbucket")
	if err != nil {
		t.Error(err)
	}
	err = cld.Upload(context.TODO(), "../test/" + filename)
	if err != nil {
		t.Error(err)
	}
}

func TestMinio(t *testing.T) {
	filename := "test-minio-file"
	createTestFile(filename)
	defer removeTestFile(filename)

	conf.Init("../conf/config.ini")

	cld, err := NewCloud(context.TODO(), "testbucket")
	if err != nil {
		if _, ok := err.(minio.BucketExistsErr); !ok {
			t.Error(err)
		}
	}
	err = cld.Upload(context.TODO(), "../test/" + filename)
	if err != nil {
		t.Error(err)
	}
}
