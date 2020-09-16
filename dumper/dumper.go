/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/9 16:07
* @Description: The file is for
***********************************************************************/

package dumper

import (
	"context"
	"errors"
	"github.com/azd1997/dumper2cloud/dumper/fakedumper"
	"github.com/azd1997/dumper2cloud/dumper/mydumper"

	"github.com/azd1997/dumper2cloud/conf"
	"github.com/azd1997/dumper2cloud/dumper/gomydumper"
)

// Dumper
type Dumper interface {
	Start() error
	Pause() error
	Continue() error
	Wait() error
}

// 新建dumper
func NewDumper(ctx context.Context, confpath string) (Dumper, error) {
	dumpertype := conf.Global().Section("dumper2cloud").Key("dumper").String()

	switch dumpertype {
	case "fakedumper":
		return fakedumper.NewFakeDumper(ctx)
	case "gomydumper":
		return gomydumper.NewGoMyDumper(ctx, confpath)
	case "mydumper":
		return mydumper.NewMyDumper(ctx)
	default:
		return nil, errors.New("unknown dumpertype")
	}
}
