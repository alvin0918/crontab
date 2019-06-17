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
	if err = master.InitWorkerMgr(); err != nil{
		goto ERT
	}

	// 初始化日志管理器
	if err = master.InitLogMgr(); err != nil {
		goto ERT
	}

	// 初始化任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERT
	}

	// 启动HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERT
	}

	ERT:
		fmt.Println(err)

}
