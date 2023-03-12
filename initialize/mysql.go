package initialize

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"minio-storage/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitMysqlDB() error {
	mysqlInfo := global.Settings.MysqlInfo

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		mysqlInfo.Username,
		mysqlInfo.Password,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.Database,
		mysqlInfo.Charset,
		mysqlInfo.ParseTime,
		mysqlInfo.Loc,
	)

	sqlDB, err := sql.Open(mysqlInfo.DriverName, args)
	if err != nil {
		log.Fatalln(err)

		return err
	}

	// 空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 打开数据库连接的最大数量，根据需求看着调
	sqlDB.SetMaxOpenConns(100)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 注册单例
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		// 禁止自动给表名加“s”
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		global.DB = nil
		log.Fatalln(err)

		return err
	}

	// 设置全局DB
	global.DB = gormDB
	log.Println("Mysql Init Success")

	return nil
}
