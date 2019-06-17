package common

import "encoding/json"

// 应答接口
type Response struct {
	Errno int `json:"errno"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

// Job 序列化
type Job struct {
	Name string `json:"name"` // 任务名称
	Command string `json:"Command"` // Shell命令
	CronExpr string `json:"cronExpr"` // cron 表达式
}

// 定义一个统一应答接口
func BuildResponse(errno int, msg string, data interface{}) (resp []byte, err error) {

	var (
		response Response
	)

	response.Data = data
	response.Msg = msg
	response.Errno = errno

	resp, err = json.Marshal(response)

	return
}
