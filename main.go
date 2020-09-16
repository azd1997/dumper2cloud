/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/9 16:06
* @Description: The file is for
***********************************************************************/

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/azd1997/dumper2cloud/conf"
	"github.com/azd1997/dumper2cloud/scheduler"
	"log"
)

var (
	// author = flag.String("author", "Eiger", "author")
	helpFlag   = flag.Bool("h", false, "display help info")
	configFlag = flag.Bool("c", false, "specify config path")
	//configPath = flag.String("config", "./conf/config.ini", "config file path")
	//cloudpath = flag.String("cloud", "", "cloud path. the real cloud path will be [CLOUDPATH]2020091009")	// 年月日时
)

func printUsage() {
	fmt.Println(`Dumper2Cloud (d2c) 
\t help dump mysql databases to cloud. now gomydumper as dumper, minio as cloud available.
\t\t USAGE:
\t\t\t d2c
\t\t\t d2c -c CONFIGPATH
\t author@Eiger`)
}

// d2c
func main() {
	flag.Parse()

	if *helpFlag {
		printUsage()
		return
	}

	configPath := "./conf/config.ini"
	if *configFlag {
		if flag.NArg() != 1 {
			log.Println("flag(-c) need arg (1)")
			printUsage()
			return
		} else {
			configPath = flag.Arg(0)
		}
	}

	// 配置初始化
	conf.Init(configPath)
	log.Printf("config init finish. configpath=%s\n", configPath)
	log.Println("config:\n",
		"dumper2cloud: ", conf.Global().Section("dumper2cloud").Keys(), "\n",
		"minio: ", conf.Global().Section("minio").Keys(), "\n",
		"mysql: ", conf.Global().Section("mysql").Keys())

	// 调度器初始化
	s, err := scheduler.NewScheduler(context.Background(), configPath)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("scheduler init finish ...")

	// 运行
	s.Run()
}
