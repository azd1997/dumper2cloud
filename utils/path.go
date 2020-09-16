/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/12/20 11:30 AM
* @Description: The file is for
***********************************************************************/

package utils

// 检查目录后缀是否有'/'，没有加上
func CheckDirSuffixSlash(dir string) string {
	n := len(dir)
	if n == 0 {
		return ""
	}

	if dir[n-1] != '/' {
		return dir + "/"
	}
	return dir
}

// 打印目录下所有文件