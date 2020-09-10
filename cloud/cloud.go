/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/9 23:50
* @Description: 云存储接口
***********************************************************************/

package cloud

import (
	"context"
	"errors"
	"github.com/azd1997/dumper2cloud/cloud/minio"
	"github.com/azd1997/dumper2cloud/conf"
)

// Cloud 云存储接口
type Cloud interface {
	// Upload 将指定文件上传
	Upload(ctx context.Context, filepath string) error
}

// NewCloud 返回新建的云存储客户端，并指定一个名称（类似于目录名，当前这次备份的所有文件都上传到该目录下）
func NewCloud(ctx context.Context, cloudpath string) (Cloud, error) {
	cloudtype := conf.Global().Section("dumper2cloud").Key("cloud").String()
	switch cloudtype {
	case "minio":
		return minio.NewMinio(ctx, cloudpath)
	default:
		return nil, errors.New("unknown cloudtype")
	}
}

