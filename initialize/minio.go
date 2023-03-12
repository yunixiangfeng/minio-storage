package initialize

import (
	"log"
	"minio-storage/global"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinioIO() error {
	minioInfo := global.Settings.MinioInfo

	// 创建minio服务，传入IP、账号、密码
	minioClient, err := minio.New(minioInfo.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(minioInfo.AccessKey, minioInfo.SecretKey, ""),
		// 关闭TLS，暂时不需要
		Secure: false,
	})

	if err != nil {
		global.MinioClient = nil
		log.Fatalln(err)

		return err
	}

	// 设置全局MinioClient
	global.MinioClient = minioClient
	log.Println("Minio Init Success")

	return nil
}
