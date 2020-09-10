/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/9/10 14:17
* @Description: 读取ini配置文件，转为全局单例。该配置单例在初始化之后只读
***********************************************************************/

package conf

import (
	"log"
	"sync"

	"github.com/go-ini/ini"
)

var cfg *ini.File

// 目前只允许初始化一次配置
var once sync.Once

// Global 获取全局配置
func Global() *ini.File {
	return cfg
}

// Init 初始化配置
func Init(confpath string) {
	once.Do(
		func() {
			c, err := loadConfig(confpath)
			if err != nil {
				log.Fatalln(err)
			}
			cfg = c
		})
}

// 加载配置
func loadConfig(confpath string) (*ini.File, error) {
	return ini.Load(confpath)
}
