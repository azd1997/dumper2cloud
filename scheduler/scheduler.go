/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 13:22
* @Description: 调度器，协调备份与上传的进度
***********************************************************************/

package scheduler

import (
	"github.com/azd1997/dumper2cloud/dumper"
	"github.com/azd1997/dumper2cloud/cloud"
)

// Scheduler 调度器
type Scheduler struct {
	dumper dumper.Dumper
	store  cloud.Storage

}

func NewScheduler(cfg )
