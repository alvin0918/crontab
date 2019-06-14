package master

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	ApiPort int `json:"apiPort"`
	ApiReadTimeout int `json:"apiReadTimeout"`
	ApiWriteTimeout int `json:"apiWriteTimeout"`
	EtcdEndpoints []string `json:"etcdEndpoints"`
	EtcdDialTimeout int `json:"etcdDialTimeout"`
	Webroot string `json:"webroot"`
	MongodbUri string `json:"mongodbUri"`
	MongodbConnectTimeout int `json:"mongodbConnectTimeout"`
}

// 配置单例
var (
	G_Config *Config
)

func InitConfig(confFile string) (err error) {

	var (
		connect []byte
		conf Config
	)

	// 读配置文件
	if connect,err = ioutil.ReadFile(confFile); err != nil {
		return err
	}

	// json反序列化
	if err = json.Unmarshal(connect, &conf); err != nil {
		return err
	}

	// 赋值单例
	G_Config = &conf

	return
}
