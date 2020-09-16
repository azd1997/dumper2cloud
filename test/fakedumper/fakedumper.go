/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/11 10:02
* @Description: 实现一个虚假的dumper，随便写文件。请确保编译结果放到bin/下
***********************************************************************/

package main

import (
	"flag"
	"github.com/azd1997/dumper2cloud/utils"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	outdir = flag.String("o", "", "output directory")
)

func main() {
	flag.Parse()

	if *outdir == "" {
		log.Fatalln("please specify output directory")
	}

	fd := &FakeDumper{
		totalsize:  100,
		singlesize: 10,
		timecost:   2,
		outputdir: *outdir,
	}

	fd.dump()
}

// FakeDumper 模拟的文件大小等价于10倍
type FakeDumper struct {
	totalsize  int // MB	总文件大小
	singlesize int // MB	单个文件大小
	timecost   int // s	写单个文件时的睡眠时间（模拟写大文件时的耗时）
	outputdir string	// 文件输出目录
}

// 模拟备份工具，向本地不断写文件，直到128M，然后就写新文件
// 由于本地测试环境资源不够，模拟备份时设定总大小100M，
// 每个文件10M， 中间给定时来模拟消耗较长时间
func (d *FakeDumper) dump() {
	text := "abcdefghi\n"
	n := len(text)

	total := 0  // 总写入数据大小
	single := 0 // 单文件写入数据大小
	filenum := 1

	file, err := os.Create(utils.CheckDirSuffixSlash(d.outputdir) + "fake-dumper-" + strconv.Itoa(filenum))
	if err != nil {
		log.Println(err)
	}

	data := "abcdefghijk\n"
	for i := 0; i < 100; i++ {
		file.WriteString(data)
	}
	for total < d.totalsize*1024*1024 {
		// 单个文件超限制
		if single > d.singlesize*1024*1024 {

			file.Close()
			time.Sleep(time.Duration(d.timecost) * time.Second)
			log.Printf("dump file [fake-dumper-%d] write finish. size %d B\n", filenum, single)

			filenum++
			file, err = os.Create(utils.CheckDirSuffixSlash(d.outputdir) + "fake-dumper-" + strconv.Itoa(filenum))
			if err != nil {
				log.Println(err)
			}
			single = 0
		}

		// 写入文件
		file.WriteString(text)
		single += n
		total += n
	}
	file.Close()
	time.Sleep(time.Duration(d.timecost) * time.Second)
	log.Printf("dump file [fake-dumper-%d] write finish. size %d B\n", filenum, single)
	log.Printf("dump finish. size %d B\n", total)

}
