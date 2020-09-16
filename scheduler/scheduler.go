/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 13:22
* @Description: 调度器，协调备份与上传的进度
***********************************************************************/

package scheduler

import (
	"context"
	"github.com/azd1997/dumper2cloud/cloud"
	"github.com/azd1997/dumper2cloud/dumper"
	"time"
)

// Scheduler 调度器
type Scheduler struct {
	Dumper dumper.Dumper
	Cloud  cloud.Cloud

	outdir string

	finishNotify chan struct{}
	dumpEndNotify chan struct{}
	uploadQueue chan string
	detectInterval int	// ms 检测间隔
}

func NewScheduler(ctx context.Context, confpath string) (*Scheduler, error) {
	outdir :=


	d, err := dumper.NewDumper(ctx, confpath)
	if err != nil {
		return nil, err
	}
	c, err := cloud.NewCloud(ctx, "")
	if err != nil {
		return nil, err
	}



	return &Scheduler{
		Dumper: d,
		Cloud:  c,
	}, nil
}

func (s *Scheduler) Run() {
	go s.detectLoop()
	go s.uploadLoop()
	go s.dumpLoop()

	<- s.finishNotify
}

func (s *Scheduler) dumpLoop() {

}

func (s *Scheduler) uploadLoop() {
	for filename := range s.uploadQueue {
		s.Cloud.Upload(context.Background(), )
	}
}

func (s *Scheduler) detectLoop() {
	ticker := time.Tick()
}