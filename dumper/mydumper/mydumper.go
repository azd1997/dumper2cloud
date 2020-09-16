/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 9:21
* @Description: The file is for
***********************************************************************/

package mydumper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/azd1997/dumper2cloud/conf"
	"os/exec"
	"syscall"
)

// NewMyDumper
func NewMyDumper(ctx context.Context) (*MyDumper, error) {
	binPath := conf.Global().Section("dumper2cloud").Key("dumper_bin_path").String()

	// 根据confpath读配置
	host := conf.Global().Section("mysql").Key("host").String()
	port := conf.Global().Section("mysql").Key("port").String()
	user := conf.Global().Section("mysql").Key("user").String()
	password := conf.Global().Section("mysql").Key("password").String()
	database := conf.Global().Section("mysql").Key("database").String()
	outdir := conf.Global().Section("mysql").Key("outdir").String()
	chunksizeMB := conf.Global().Section("mysql").Key("chunksize").String()

	cmd := exec.CommandContext(ctx, binPath, "-h", host, "-P", port, "-u", user, "-p", password,
		"-B", database, "-d", outdir, "-F", chunksizeMB)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	return &MyDumper{cmd: cmd, out: out}, nil
}

type MyDumper struct {
	cmd *exec.Cmd
	out bytes.Buffer
}

func (d *MyDumper) Start() error {
	if d.cmd == nil {
		return errors.New("d.cmd is nil")
	}
	return d.cmd.Start()
}

func (d *MyDumper) Pause() error {
	if d.cmd == nil {
		return errors.New("d.cmd is nil")
	}
	if d.cmd.Process == nil {
		return errors.New("d.cmd is not started")
	}
	return d.cmd.Process.Signal(syscall.SIGSTOP)
}

func (d *MyDumper) Continue() error {
	if d.cmd == nil {
		return errors.New("d.cmd is nil")
	}
	if d.cmd.Process == nil {
		return errors.New("d.cmd is not started")
	}
	return d.cmd.Process.Signal(syscall.SIGCONT)
}

func (d *MyDumper) Wait() error {
	if d.cmd == nil {
		return errors.New("d.cmd is nil")
	}
	if d.cmd.Process == nil {
		return errors.New("d.cmd is not started")
	}

	_, err := d.cmd.Process.Wait()
	if err != nil {
		return err
	}

	// 输出cmd的执行输出
	fmt.Println()
	fmt.Printf("d.cmd [%s] Stdout/Stderr: \n", d.cmd.String())
	if d.out.Len() == 0 {
		fmt.Println("everything seems ok.")
	} else {
		fmt.Println(d.out.String())
	}
	fmt.Println()

	return nil
}
