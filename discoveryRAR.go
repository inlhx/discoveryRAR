package main

/**
* https://github.com/inlhx   (https://gitee.com/aliu)
**/
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

var passpath string = "password.txt"
var filepath string = ""
var storeDir string = "unrar-path"

func main() {
	fmt.Println("RAR压缩文件探索器,根据密码字典查找您丢失的密码")
	fmt.Println("Rar compressed file explorer to find your lost password according to the password dictionary")
	fmt.Println("本软件基于MIT开源,仅供学习交流,切勿用于非法用途! ")
	fmt.Println("This software is based on MIT open source and is only for learning. Do not use it for illegal purposes!! ")
	fmt.Println("Create By https://github.com/inlhx   (https://gitee.com/aliu)  ")
	fmt.Printf("\r\n\r\n\r\n\r\n请输入需要搜寻的文件全路径  \r\n Example : c:\\logs\\example.rar or example.rar")
	fmt.Printf("\r\n\r\n请输入文件路径:")
	fmt.Scanf("%v %v\n", &filepath)
	_, err1 := os.Lstat(filepath)
	if err1 != nil {
		if os.IsNotExist(err1) {
			fmt.Printf("您输入的" + filepath + "文件没找到,请检查")
			time.Sleep(time.Duration(300) * time.Second)
			os.Exit(0)
		}
	}

	_, err2 := os.Lstat(passpath)
	if os.IsNotExist(err2) {
		fmt.Printf("\r\n\r\n\r\n\r\n\r\n\r\n错误:没找到密码本\r\n请检查程序目录下是否有" + passpath + "文件存在,如果不存在请使用generate-password.exe生成密码本")
		time.Sleep(time.Duration(20) * time.Second)
		os.Exit(0)
	}

	err := os.MkdirAll(storeDir, 0755)
	if err != nil {
		fmt.Printf("\r\n\r\n\r\n\r\n无法创建storeDir 文件夹,请手动创建" + storeDir + "文件夹")
		time.Sleep(time.Duration(300) * time.Second)
		os.Exit(0)
	}
	readPassword(passpath)
}

func unRarCmd(rarpath string, pass string) {
	fmt.Println(pass)
	cmd := exec.Command("storage.data", "e", "-p"+pass, rarpath, "unrar-path") //文件解压保存到unrar-path
	out, _ := cmd.Output()
	// var outLen = len(out)
	var lastLen = out[len(out)-3]
	// fmt.Println("检查正确信息长度:", outLen, out)
	// fmt.Println("xxxxxxxxx:", out[outLen-3])
	if lastLen == 100 || lastLen == 75 { //正确信息长度,
		fmt.Printf("密码为：%s \n", pass)
	//	for {
	//		time.Sleep(time.Duration(10) * time.Second)
	//	}
		os.Exit(0)
	}
}

func readPassword(passpath string) {
	fp, _ := os.OpenFile(passpath, os.O_RDONLY, 6)
	defer fp.Close()

	// 创建文件的缓存区
	r := bufio.NewReader(fp)
	for {
		pass, err2 := r.ReadBytes('\n')
		if err2 == io.EOF {
			break
		}
		pass = pass[:len(pass)-2] // 去除末尾 /n
		time.Sleep(time.Duration(15) * time.Millisecond) //延迟避免堵塞
		go unRarCmd(filepath, string(pass))
	}
}
