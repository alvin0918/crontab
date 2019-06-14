package main

import (
	"flag"
	"runtime"
	"github.com/alvin0918/crontab/master"
	"fmt"
)

var confFile string

/**
	初始化命令行参数
 */
func initArgs() {
	flag.StringVar(&confFile, "c", "./master.json", "指定配置文件")
	flag.Parse()
}

/**
	初始化线程数
 */
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main()  {

	var (
		err error
	)

	// 初始化命令行参数
	initArgs()

	// 初始化线程数
	initEnv()

	// 读取配置文件
	if err = master.InitConfig(confFile); err != nil {
		goto ERT
	}

	// 初始化服务发现模块



	ERT:
		fmt.Println(err)

}
