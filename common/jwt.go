package common

import (
	"minio-storage/models"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("yonu_key")

type Claims struct {
	UserID    int16
	Access    string
	AccessKey string
	Level     int
	jwt.StandardClaims
}

// 颁发token
func ReleaseToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserID:    user.UserID,    //用户id
		Access:    user.Access,    // 用户权限
		AccessKey: user.AccessKey, // 用户账户
		Level:     user.Level,     // 等级
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),     // 签发时间
			Issuer:    "minio",               // 签发人
			Subject:   "token",               // 标题
		},
	}

	// 加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
