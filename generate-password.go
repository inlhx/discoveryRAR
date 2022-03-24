package main

/**
* https://github.com/inlhx   (https://gitee.com/aliu)
**/
import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"

	"os"
)

var (
	datapath        = "./password.txt"
	imark    uint32 = 0
)

func gen(charset string, n int, sc chan string) {

	for i, c := range charset {
		if n == 1 {
			sc <- string(c)
		} else {
			var ssc = make(chan string)
			go gen(charset[:i]+charset[i+1:], n-1, ssc)
			for k := range ssc {
				sc <- fmt.Sprintf("%v%v", string(c), k)
			}
		}
	}
	close(sc)

}

func main() {
	var minlen = 0
	var maxlen = 0
	var model = 0
	fmt.Println("密码生成器,递归生成组合密码")
	fmt.Println("本软件基于MIT开源,仅供学习交流,切勿用于非法用途! ")
	fmt.Println("This software is based on MIT open source and is only for learning. Do not use it for illegal purposes!! ")
	fmt.Println("Create By https://github.com/inlhx   (https://gitee.com/aliu)")
	fmt.Println("长度越长,复杂度越高,生成时间越长文件越大,3-8位长度大概有5GB")

	var model1 = "0123456789"
	var model2 = "0123456789abcdefghijklmnopqrstuvwxyz"
	var model3 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var model4 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+`-=[]\\{}|;':\",./<>? "
	var modelSelect = ""

	fmt.Printf("\r\n\r\n\r\n\r\n请输入密码最小长度:")
	fmt.Scanf("%v %v\n", &minlen)
	fmt.Printf("请输入密码最大长度:")
	fmt.Scanf("%v %v\n", &maxlen)
	fmt.Printf("\r\n\r\n\r\n\r\n请选择密码组合模式: \r\n1.纯数字 " + model1 + "  \r\n2.数字+字母小写 " + model2 + " \r\n3.数字+字母大小写 " + model3 + "  \r\n 4.数字+字母大小写+字符混合  " + model4)
	fmt.Printf("\r\n\r\n请选择密码组合模式1,2,3,4:")
	fmt.Scanf("%v %v\n", &model)
	maxlen = maxlen + 1
	if model == 1 {
		modelSelect = model1
	} else if model == 2 {
		modelSelect = model2
	} else if model == 3 {
		modelSelect = model3
	} else if model == 4 {
		modelSelect = model4
	}
	fmt.Println()
	starttime := time.Now()

	fmt.Println("开始时间:", starttime)
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := minlen; i < maxlen; i++ {
		sc := make(chan string)

		fmt.Println("正在生成", i, "位长度密码", time.Now())
		go gen(modelSelect, i, sc)

		fs, e := os.OpenFile(datapath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if e != nil {
			panic(e)
		}
		defer fs.Close()

		for x := range sc {
			atomic.AddUint32(&imark, 1)
			fs.WriteString(x)
			fs.WriteString("\r\n")
			// fmt.Println("Gen:", x)
		}
	}
	imarkFinal := atomic.LoadUint32(&imark)
	since := int(time.Since(starttime).Seconds())

	fmt.Println("结束时间:", time.Now())
	fmt.Println("完成,耗时:", since, "s", "生成:", imarkFinal, "个密码,密码已写入password.txt")
	time.Sleep(time.Duration(300) * time.Second)
}
