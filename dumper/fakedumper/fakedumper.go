/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/11 10:02
* @Description: 针对fakedumper（一个假的备份工具），而实现的dumper封装
***********************************************************************/

package main

import (
	"context"
	"fmt"
	"github.com/azd1997/dumper2cloud/conf"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)


// NewFakeDumper
func NewFakeDumper(ctx context.Context) (*FakeDumper, error) {
	binPath := conf.Global().Section("dumper2cloud").Key("dumper_bin_path").String()
	cmd := exec.CommandContext(ctx, binPath)

	return &FakeDumper{cmd:cmd}, nil
}

type FakeDumper struct {
	cmd *exec.Cmd
}

func (d *FakeDumper) Start() error {
	return d.cmd.Start()
}

func (d *FakeDumper) Pause() error {
	d.cmd.Process.Signal(syscall.SIGHUP)
	return d.cmd.Start()
}

func getPID(appName string) int {
	cmd := exec.Command("ps", "-C", appName)
	output, _ := cmd.Output()

	fields := strings.Fields(string(output))

	for _, v := range fields {
		if v == appName {
			return true
		}
	}

	return false
}
