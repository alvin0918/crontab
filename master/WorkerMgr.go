package master

import (
	"go.etcd.io/etcd/clientv3"
	"time"
)

type WorkerMgr struct {
	kv clientv3.KV
	lease clientv3.Lease
	client *clientv3.Client
}

var (
	G_WorkerMgr *WorkerMgr
)

/**
  初始化服务发现
 */
func InitWorkerMgr() (err error) {

	var (
		config clientv3.Config
		client *clientv3.Client
		kv clientv3.KV
		lease clientv3.Lease
	)

	// 初始化配置
	config = clientv3.Config{
		Endpoints: G_Config.EtcdEndpoints, // 集群地址
		DialTimeout: time.Duration(G_Config.EtcdDialTimeout) * time.Millisecond,  // 超时时间
	}

	// 建立连接
	if client,err = clientv3.New(config); err != nil {
		return err
	}

	// 得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	G_WorkerMgr = &WorkerMgr{
		kv:kv,
		lease:lease,
		client:client,
	}

	return
}
