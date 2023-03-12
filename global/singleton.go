package global

import (
	"minio-storage/conf"

	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

var (
	Settings    conf.ServerConf
	DB          *gorm.DB
	MinioClient *minio.Client
)
