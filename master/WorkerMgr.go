package master

import (
	"go.etcd.io/etcd/clientv3"
	"time"
)

func InitWorkerMgr() (err error) {

	var (
		config clientv3.Config
	)

	// 初始化配置
	config = clientv3.Config{
		Endpoints: G_Config.EtcdEndpoints, // 集群地址
		DialTimeout: time.Duration(G_Config.EtcdDialTimeout) * time.Millisecond,  // 超时时间
	}

	

	return
}
