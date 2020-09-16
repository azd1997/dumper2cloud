/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 9:21
* @Description: The file is for
***********************************************************************/

package gomydumper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/azd1997/dumper2cloud/conf"
	"os/exec"
	"syscall"
)

// NewGoMyDumper
func NewGoMyDumper(ctx context.Context, confpath string) (*GoMyDumper, error) {
	binPath := conf.Global().Section("dumper2cloud").Key("dumper_bin_path").String()
	cmd := exec.CommandContext(ctx, binPath, "-c", confpath)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	return &GoMyDumper{cmd:cmd, out:out}, nil
}

type GoMyDumper struct {
	cmd *exec.Cmd
	out bytes.Buffer
}

func (d *GoMyDumper) Start() error {
	if d.cmd == nil {
		return errors.New("d.cmd is nil")
	}
	return d.cmd.Start()
}

func (d *GoMyDumper) Pause() error {
	if d.cmd == nil {
		return errors.New("d.cmd is nil")
	}
	if d.cmd.Process == nil {
		return errors.New("d.cmd is not started")
	}
	return d.cmd.Process.Signal(syscall.SIGSTOP)
}

func (d *GoMyDumper) Continue() error {
	if d.cmd == nil {
		return errors.New("d.cmd is nil")
	}
	if d.cmd.Process == nil {
		return errors.New("d.cmd is not started")
	}
	return d.cmd.Process.Signal(syscall.SIGCONT)
}

func (d *GoMyDumper) Wait() error {
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
