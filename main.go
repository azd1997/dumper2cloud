/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/9 16:06
* @Description: The file is for
***********************************************************************/

package main

import (
	"flag"
	"fmt"
	"github.com/azd1997/dumper2cloud/conf"
	"github.com/azd1997/dumper2cloud/scheduler"
)

var (
	// author = flag.String("author", "Eiger", "author")
	help = flag.Bool("help", false, "display help info")
	configPath = flag.String("config", "./conf/config.ini", "config file path")
	cloudpath = flag.String("cloud", "", "cloud path. the real cloud path will be [CLOUDPATH]2020091009")	// 年月日时
)

func printUsage() {
	fmt.Println(`Dumper2Cloud (d2c) 
\t help dump mysql databases to cloud. now gomydumper as dumper, minio as cloud available.
\t\t USAGE: d2c -c CONFIGPATH
\t author@Eiger`)
}

// d2c
func main() {
	flag.Parse()
	if *help {
		printUsage()
		return
	}

	// 配置初始化
	conf.Init(*configPath)

	// 调度器初始化
	s := scheduler.NewScheduler()
	s.Start()
}
