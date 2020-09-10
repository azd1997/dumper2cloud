/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 9:21
* @Description: The file is for
***********************************************************************/

package mydumper

import (
	"context"
	"os/exec"
)

// NewMyDumper 新构建一个MyDumper工具
func NewMyDumper(bin string, args ...string) (*MyDumper, error) {
	// 检查bin有效性
	// 检查bin作为路径时，对应的是不是文件，文件是不是可执行
	// 如果bin没有"/"，使用exec.LoopPath尝试去PATH目录下找

	return &MyDumper{bin:bin}, nil
}

type MyDumper struct {
	bin string
}

func (d *MyDumper) Dump() {
	cmd := exec.CommandContext(context.TODO(), d.bin, )
	cmd.CombinedOutput()
}
