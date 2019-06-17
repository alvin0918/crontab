package master

import (
	"net/http"
	"github.com/alvin0918/crontab/common"
	"encoding/json"
)

/**
  设置JOB
    POST job={"name": "job1", "command": "echo hello", "cronExpr": "* * * * *"}
 */
func handleJobSave(resp http.ResponseWriter, req *http.Request)  {

	var (
		err error
		bytes []byte
		postJob string
		job common.Job
		oldJob *common.Job
	)

	// 获取POST配置
	if err = req.ParseForm(); err != nil {
		goto ERT
	}

	// 获取job字段
	postJob = req.PostForm.Get("job")

	// 序列化JOB
	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERT
	}

	// 将任务保存到ETCD
	if oldJob, err = G_jobMgr.SaveJob(&job); err != nil {
		return
	}

	// 返回正常应答
	if bytes, err = common.BuildResponse(0, "success", oldJob); err != nil {
		resp.Write(bytes)
	}

	return

	ERT:
		bytes, err = common.BuildResponse(-1, err.Error(), nil)
		resp.Write(bytes);
}

/**
	删除任务
		POST /job/delete   name=job1
 */
func handleJobDelete(resp http.ResponseWriter, req *http.Request) {

	var (
		err error
		bytes []byte
		name string
		oldJob *common.Job
	)

	// 获取参数
	if err = req.ParseForm(); err != nil {
		goto ERT
	}

	// 获取需要删除任务名
	name = req.PostForm.Get("name")

	// 删除ETCD中的任务
	if oldJob, err = G_jobMgr.DeleteJob(name); err != nil {
		return
	}

	// 返回应答
	if bytes, err = common.BuildResponse(0, "success", oldJob); err != nil {
		resp.Write(bytes)
	}

	ERT:
		bytes, err = common.BuildResponse(-1, err.Error(), nil)
		resp.Write(bytes);
}

func handleJobList(resp http.ResponseWriter, req *http.Request) {

}

func handleJobKill(resp http.ResponseWriter, req *http.Request) {

}

func handleJobLog(resp http.ResponseWriter, req *http.Request) {

}

func handleWorkerList(resp http.ResponseWriter, req *http.Request) {

}

func InitApiServer() (err error) {

	var (
		mux *http.ServeMux
	)

	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)
	mux.HandleFunc("/job/delete", handleJobDelete)
	mux.HandleFunc("/job/list", handleJobList)
	mux.HandleFunc("/job/kill", handleJobKill)
	mux.HandleFunc("/job/log", handleJobLog)
	mux.HandleFunc("/worker/list", handleWorkerList)
	// 启动HTTP服务

	return
}
