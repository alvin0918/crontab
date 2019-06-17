package master

import (
	"go.etcd.io/etcd/clientv3"
	"github.com/alvin0918/crontab/common"
	"encoding/json"
	"golang.org/x/net/context"
	"time"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

type JobMgr struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
}

var (
	// 单例
	G_jobMgr *JobMgr
)

func InitJobMgr() (err error) {
	var (
		config clientv3.Config
		client *clientv3.Client
		kv clientv3.KV
		lease clientv3.Lease
	)

	// 初始化配置
	config = clientv3.Config{
		Endpoints: G_Config.EtcdEndpoints, // 集群地址
		DialTimeout: time.Duration(G_Config.EtcdDialTimeout) * time.Millisecond, // 连接超时
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		return
	}

	// 得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	G_jobMgr = &JobMgr{
		client :client,
		kv: kv,
		lease: lease,
	}
	return
}

func (jobMgr *JobMgr) SaveJob(job *common.Job) (oldJob *common.Job, err error) {

	var (
		jobKey string
		jobValue []byte
		putResp *clientv3.PutResponse
		oldJobObj common.Job
	)

	// 保存到etct中的KEY
	jobKey = common.JOB_SAVE_DIR + job.Name

	// 任务JSON
	if jobValue, err = json.Marshal(job); err != nil {
		return
	}

	// 保存到ETCD
	if putResp, err = jobMgr.kv.Put(context.TODO(), jobKey, string(jobValue), clientv3.WithPrevKV()); err != nil {
		return
	}

	// 如果是更新，那么返回旧值
	if putResp.PrevKv != nil {
		// 对旧值序列化
		if err = json.Unmarshal(putResp.PrevKv.Value, &oldJobObj); err != nil {
			err = nil
			return
		}
	}

	oldJob = &oldJobObj

	return
}

// 删除任务
func (JobMgr *JobMgr) DeleteJob(name string) (oldJob *common.Job, err error)  {

	var (
		jobKey string
		delResp *clientv3.DeleteResponse
		oldJobObj common.Job
	)
	
	// 构建ETCD KEY
	jobKey = common.JOB_SAVE_DIR + name
	
	if delResp, err = JobMgr.kv.Delete(context.TODO(), jobKey, clientv3.WithPrevKV()); err != nil {
		return 
	}

	if len(delResp.PrevKvs) != 0 {
		// 解析旧值
		if err = json.Unmarshal(delResp.PrevKvs[0].Value, &oldJobObj); err != nil {
			err = nil
			return
		}

		oldJob = &oldJobObj
	}
	
	return 
}

// 获取任务列表
func (jobMgr *JobMgr) ListJobs() (jobList []*common.Job, err error) {

	var (
		dirKey string
		getResp *clientv3.GetResponse
		kvPair *mvccpb.KeyValue
		job *common.Job
	)

	// 获取任务保存列表
	dirKey = common.JOB_SAVE_DIR

	// 获取目录下所有任务
	if getResp, err = jobMgr.kv.Get(context.TODO(), dirKey, clientv3.WithPrevKV()); err != nil {
		return
	}

	// 初始化数组空间
	jobList = make([]*common.Job, 0)

	// 遍历所有的任务，进行反序列化
	for _, kvPair = range getResp.Kvs{
		job = &common.Job{}
		if err = json.Unmarshal(kvPair.Value, job); err != nil {
			err = nil
			continue
		}
		jobList = append(jobList, job)
	}

	return
}

// 强制杀死任务
func (jobMgr *JobMgr) KillJob(name string) (err error) {

	var (
		killKey string
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
	)

	killKey = common.JOB_KILLER_DIR + name

	if leaseGrantResp, err = jobMgr.lease.Grant(context.TODO(), 1);err != nil {
		return
	}

	leaseId = leaseGrantResp.ID

	if _, err = jobMgr.kv.Put(context.TODO(), killKey, "", clientv3.WithLease(leaseId)); err != nil {
		return
	}

	return
}



































