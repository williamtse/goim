package auth

import (
	"GoIM/pkg/common/config"
	"GoIM/pkg/common/constant"
	"GoIM/pkg/common/db"
	"GoIM/pkg/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"strconv"
	"time"
)

type Auth struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
}

func (a *Auth) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *Auth) UnmarshalBinary(data []byte) (err error) {
	return json.Unmarshal(data, a)
}

func getTokenKey(token string) (key string) {
	return constant.AccessTokenPrefix + token
}

func getUidKey(uid string) string {
	return constant.UidPrefix + uid
}

func GetUserAuthInfoByUid(uid string) *Auth {
	userAuth := &Auth{}
	if err := db.DB.RDB.Get(context.Background(), getUidKey(uid)).Scan(userAuth); err != nil {
		return nil
	}
	return userAuth
}

func GetUserAuthInfo(token string) (auth *Auth) {
	userAuth := &Auth{}
	if err := db.DB.RDB.Get(context.Background(), getTokenKey(token)).Scan(userAuth); err != nil {
		return nil
	}
	return userAuth
}

func checkToken(token string) bool {
	userAuth := GetUserAuthInfo(token)
	if userAuth != nil {
		log.Println(userAuth)
		return true
	}
	return false
}

func AccessTokenMiddleware(netxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		tokenStr := values.Get("token")
		if checkToken(tokenStr) {
			netxt.ServeHTTP(w, r)
		} else {
			log.Fatalf("access deny!")
		}
	})
}

func GetKey(str string) (key string) {
	return config.Config.Secret + str
}

func Login(appId string, appSecret string, uid string) (token string, err error) {
	now := time.Now().Unix()
	tokenStr := strconv.FormatInt(now, 10)
	idInt, nil := strconv.ParseInt(uid, 10, 32)
	authUser := &Auth{
		Id:          idInt,
		AccessToken: tokenStr,
	}
	db.DB.RDB.Set(context.Background(), getTokenKey(tokenStr), authUser, 0)
	db.DB.RDB.Set(context.Background(), getUidKey(uid), authUser, 0)
	return tokenStr, nil
}

func Logout(token string) error {
	authUser := GetUserAuthInfo(token)
	if authUser != nil {
		db.DB.RDB.Del(context.Background(), getTokenKey(authUser.AccessToken))
		db.DB.RDB.Del(context.Background(), getUidKey(utils.Int64ToString(authUser.Id)))
	}
	return nil
}
