package auth

import (
	"GoIM/pkg/common/config"
	"time"

	jwtPkg "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	ID       uint64 `json:"user_id"`
	Username string `json:"username"`
	jwtPkg.StandardClaims
}

const TokenExpireDuration = time.Hour * 24 * 2 // 过期时间 -2天
var Secret = []byte("i am zhaohaiyu")          // Secret(盐) 用来加密解密

func GenerateToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "frank"
	claims := JwtClaims{
		ID:       10001,
		Username: "frank",
		StandardClaims: jwtPkg.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims).SignedString([]byte(config.Config.Secret))
	return token, err
}
