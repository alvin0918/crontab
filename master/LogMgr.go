package master

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/net/context"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
)

// mongodb日志管理
type LogMgr struct {
	client *mongo.Client
	logCollection *mongo.Collection
}

var (
	G_LogMgr *LogMgr
)

func InitLogMgr() (err error) {

	var (
		client *mongo.Client
	)

	// 建立mongodb连接
	if client, err = mongo.Connect(
		context.TODO(),
		G_Config.MongodbUri,
		clientopt.ConnectTimeout(time.Duration(G_Config.MongodbConnectTimeout) * time.Millisecond)); err != nil {
			return err
	}

	G_LogMgr = &LogMgr{
		client:client,
		logCollection: client.Database("cron").Collection("log"),
	}

	return
}

