/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/9 16:07
* @Description: The file is for
***********************************************************************/

package dumper

import (
	"context"
	"errors"

	"github.com/azd1997/dumper2cloud/conf"
	"github.com/azd1997/dumper2cloud/dumper/gomydumper"
)

// Dumper
type Dumper interface {
	Dump(ctx context.Context) error
}

func NewDumper(ctx context.Context) (Dumper, error) {
	dumpertype := conf.Global().Section("dumper2cloud").Key("dumper").String()

	switch dumpertype {
	case "gomydumper":
		return gomydumper.NewGomydumper(ctx, bin, conf)
	default:
		return nil, errors.New("unknown dumpertype")
	}
}
