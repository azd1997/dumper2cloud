/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 9:21
* @Description: The file is for
***********************************************************************/

package gomydumper

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// NewGomydumper 新构建一个go-mydumper工具
func NewGomydumper(bin string, conf string) (*Gomydumper, error) {
	// 检查bin有效性
	// 检查bin作为路径时，对应的是不是文件，文件是不是可执行
	// 如果bin没有"/"，使用exec.LoopPath尝试去PATH目录下找

	return &Gomydumper{bin:bin, conf:conf}, nil
}

type Gomydumper struct {
	bin string
	conf string
}

func (d *Gomydumper) Dump(ctx context.Context) error {
	// 直接调用
	cmd := exec.CommandContext(ctx, d.bin, )
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(out))

	cmd.Process.Wait()

	// 查看bin进程ID
	//
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
