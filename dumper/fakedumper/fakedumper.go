/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/11 10:02
* @Description: 针对fakedumper（一个假的备份工具），而实现的dumper封装
***********************************************************************/

package fakedumper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/azd1997/dumper2cloud/conf"
)

// NewFakeDumper
func NewFakeDumper(ctx context.Context) (*FakeDumper, error) {
	binPath := conf.Global().Section("dumper2cloud").Key("dumper_bin_path").String()
	cmd := exec.CommandContext(ctx, binPath)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	return &FakeDumper{cmd: cmd, out: out}, nil
}

type FakeDumper struct {
	cmd *exec.Cmd
	out bytes.Buffer
}

func (d *FakeDumper) Start() error {
	if d.cmd == nil {
		return errors.New("d.cmd is nil")
	}
	return d.cmd.Start()
}

func (d *FakeDumper) Pause() error {
	if d.cmd == nil {
		return errors.New("d.cmd is nil")
	}
	if d.cmd.Process == nil {
		return errors.New("d.cmd is not started")
	}
	return d.cmd.Process.Signal(syscall.SIGSTOP)
}

func (d *FakeDumper) Continue() error {
	if d.cmd == nil {
		return errors.New("d.cmd is nil")
	}
	if d.cmd.Process == nil {
		return errors.New("d.cmd is not started")
	}
	return d.cmd.Process.Signal(syscall.SIGCONT)
}

func (d *FakeDumper) Wait() error {
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
	fmt.Printf("d.cmd [%s] Stdout/Stderr: \n", d.cmd.String())
	fmt.Println(d.out.String())
	return nil
}
