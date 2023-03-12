package initialize

import (
	"io/ioutil"
	"log"
	"minio-storage/conf"
	"minio-storage/global"
	"os"

	"gopkg.in/yaml.v3"
)

func InitConfig() error {
	workDir, _ := os.Getwd()
	// 读取yaml配置文件
	yamlFile, err := ioutil.ReadFile(workDir + "/conf/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err %v", err)
		return err
	}

	// 配置信息模型
	serverConfig := conf.ServerConf{}

	// 将yaml文件对应配置信息写入serverConfig
	err = yaml.Unmarshal(yamlFile, &serverConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}

	// 设置全局Settings
	global.Settings = serverConfig

	return nil
}
