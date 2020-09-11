/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 13:22
* @Description: 调度器，协调备份与上传的进度
***********************************************************************/

package scheduler

import (
	"context"
	"github.com/azd1997/dumper2cloud/dumper"
	"github.com/azd1997/dumper2cloud/cloud"
)

// Scheduler 调度器
type Scheduler struct {
	Dumper dumper.Dumper
	Cloud  cloud.Cloud

}

func NewScheduler(ctx context.Context) (*Scheduler, error) {
	// 生成当前时间，加上

	d, err := dumper.NewDumper(ctx)
	if err != nil {
		return nil, err
	}
	c, err := cloud.NewCloud(ctx, "")
	if err != nil {
		return nil, err
	}
}
