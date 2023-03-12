package repositories

import (
	"errors"
	"minio-storage/common"
	"minio-storage/global"
	"minio-storage/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Login(*models.User) (string, error)
}

type UserManageRepository struct {
	table string
}

func (*UserManageRepository) Login(user *models.User) (string, error) {
	if global.DB == nil {
		return "", errors.New("数据库连接失败")
	}

	// 初始化user表，不是必须
	global.DB.AutoMigrate(&models.User{})

	var m models.User

	// 判断用户是否存在
	if err := global.DB.Where("access_key=?", &user.AccessKey).First(&m).Error; err != nil {
		if m.UserID == 0 {
			return "", errors.New("用户不存在")
		}
		return "", err
	}

	// 数据库的密码没有用hash加密的话，就不需要通过bcrypt库的方法来对比，直接对比就好
	// bcrypt库的方法同样可以加密后插入数据库
	if err := bcrypt.CompareHashAndPassword([]byte(m.SecretKey), []byte(user.SecretKey)); err != nil {
		return "", errors.New("密码错误")
	}

	// 颁发token
	token, err := common.ReleaseToken(m)
	if err != nil {
		return "", err
	}

	return token, nil
}
