/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 13:22
* @Description: 调度器，协调备份与上传的进度
***********************************************************************/

package scheduler

import (
	"context"
	"fmt"
	"github.com/azd1997/dumper2cloud/cloud"
	"github.com/azd1997/dumper2cloud/conf"
	"github.com/azd1997/dumper2cloud/dumper"
	"github.com/azd1997/dumper2cloud/utils"
	"log"
	"os"
	"path/filepath"
	"time"
)

// block 用来表示一个已经准备好上传的文件。为了实现失败重试，加入了failnum字段，失败超过3次后就报错放弃
type block struct {
	path    string
	failnum int
}

// Scheduler 调度器
type Scheduler struct {
	Dumper dumper.Dumper
	Cloud  cloud.Cloud

	outdir string

	finishNotify  chan struct{}
	dumpEndNotify chan struct{}
	uploadQueue   chan *block

	detectInterval int // ms 检测间隔

	detected map[string]bool // 已检测到可以上传的文件名构成的表

	maxSpace    int
	warnSpace   int64
	resumeSpace int64
	//warnFactor float64	// 当备份文件总大小 > maxspace * warnfactor 时，通知dumper进程暂停
	//resumeFactor float64 // 当备份文件本地总占用 < maxspace * resumeFactor 时， 通知dumper恢复
}

func NewScheduler(ctx context.Context, confpath string) (*Scheduler, error) {

	// dumper
	d, err := dumper.NewDumper(ctx, confpath)
	if err != nil {
		return nil, err
	}

	// cloud
	now := time.Now()
	cldpath := fmt.Sprintf("%s-%d%d%d-%d%d", cloudpath(),
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	c, err := cloud.NewCloud(ctx, cldpath)
	if err != nil {
		return nil, err
	}

	// maxspace
	maxspace := maxSpace()

	return &Scheduler{
		Dumper: d,
		Cloud:  c,

		outdir: utils.CheckDirSuffixSlash(outdir()),

		finishNotify:  make(chan struct{}),
		dumpEndNotify: make(chan struct{}),
		uploadQueue:   make(chan *block, uploadQueueSize()),

		detectInterval: detectInterval(),
		detected:       make(map[string]bool),

		maxSpace:    maxspace,
		warnSpace:   warnSpace(maxspace),
		resumeSpace: resumeSpace(maxspace),
	}, nil
}

func (s *Scheduler) Run() {
	go s.detectLoop()
	go s.uploadLoop()
	go s.dumpLoop()

	<-s.finishNotify
	log.Println("dumper2cloud finish ...")
}

func (s *Scheduler) dumpLoop() {
	// 开始备份
	err := s.Dumper.Start()
	if err != nil {
		log.Fatalln(err)
	}

	// dumper进程结束
	err = s.Dumper.Wait()
	if err != nil {
		log.Println(err)
	}

	log.Println("dumpLoop quit ...")

	// 通知
	close(s.dumpEndNotify)
}

func (s *Scheduler) uploadLoop() {
	// 接收uploadQueue传来的path，并上传，直至uploadQueue被关闭
	for block := range s.uploadQueue {
		b := block

		err := s.Cloud.Upload(context.Background(), b.path)
		if err != nil { // 上传失败，考虑重试
			log.Printf("upload %s error: %s\n", b.path, err)
			b.failnum++
			if b.failnum >= 3 {
				log.Printf("upload %s fail 3 times, dropped.\n", b.path)
			} else {
				s.uploadQueue <- b
			}
		} else { // 上传成功，将本地文件删除
			log.Printf("upload %s succ, delete it now\n", b.path)
			os.Remove(b.path)
		}
	}
	log.Println("uploadLoop quit ...")
	close(s.finishNotify)
}

func (s *Scheduler) detectLoop() {
	ticker := time.Tick(time.Duration(s.detectInterval) * time.Millisecond)

	for {
		select {
		case <-s.dumpEndNotify:
			// 把所有剩余的文件都塞入uploadQueue
			s.walkOutdir()
			// 关闭uploadQueue并退出
			close(s.uploadQueue)

			log.Println("detectLoop quit ...")
			return
		case <-ticker:
			// 检查当前目录的情况
			stat, err := os.Stat(s.outdir)
			if err != nil {
				log.Fatalf("stat outdir error: %s\n", err)
			}
			mb := stat.Size() >> 20
			if mb > s.warnSpace {
				s.Dumper.Pause() // 暂停dumper进程
			} else if mb < s.resumeSpace {
				s.Dumper.Continue() // 恢复dumper进程
			}
			// 扫描outdir，把备份好的文件交给uploadLoop上传
			s.walkOutdir()
		}
	}
}

// 遍历outdir，把符合条件的文件名传给uploadLoop去上传
func (s *Scheduler) walkOutdir() {
	filepath.Walk(s.outdir, func(path string, info os.FileInfo, err error) error {
		// 是否是文件?
		if info.IsDir() {
			return nil
		}
		// 是否已检测过
		if s.detected[info.Name()] {
			return nil
		}
		// 是否已写完，没有被其他进程占用
		if _, err := os.Open(path); err != nil {
			return nil
		}
		// 加入uploadQueue
		s.uploadQueue <- &block{path: path}
		// 加入detected
		s.detected[path] = true

		return nil
	})
}

//////////////////  配置   ////////////////

var outdir = func() string {
	return conf.Global().Section("mysql").Key("outdir").String()
}

var detectInterval = func() int {
	i, err := conf.Global().Section("dumper2cloud").Key("detect_interval").Int()
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

var cloudpath = func() string {
	return conf.Global().Section("dumper2cloud").Key("cloud_path").String()
}

var uploadQueueSize = func() int {
	i, err := conf.Global().Section("dumper2cloud").Key("upload_queue_size").Int()
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

var maxSpace = func() int {
	i, err := conf.Global().Section("dumper2cloud").Key("max_space").Int()
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

var warnSpace = func(max int) int64 {
	i, err := conf.Global().Section("dumper2cloud").Key("warn_factor").Float64()
	if err != nil {
		log.Fatalln(err)
	}
	return int64(i * float64(max))
}

var resumeSpace = func(max int) int64 {
	i, err := conf.Global().Section("dumper2cloud").Key("resume_factor").Float64()
	if err != nil {
		log.Fatalln(err)
	}
	return int64(i * float64(max))
}
