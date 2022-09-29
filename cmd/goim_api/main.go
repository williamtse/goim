package main

import (
	"GoIM/pkg/auth"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 解决跨域问题
func Core() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "True")
		//放行索引options
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		//处理请求
		c.Next()
	}
}

// 登录获取token
func doLogin(c *gin.Context) {
	log.Println("do login")
	name := c.PostForm("username")
	id := c.PostForm("id")
	token := time.Now().Unix()
	tokenStr := strconv.FormatInt(token, 10)
	key := tokenStr + name
	idInt, nil := strconv.ParseInt(id, 10, 32)
	auth := &auth.Auth{
		Id:   idInt,
		Name: name,
	}
	cacheKey := auth.GetKey(key)
	if err2 := auth.RedisCli.Set(cacheKey, auth, 0).Err(); err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"code": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": cacheKey,
			"code": 1,
		})
	}
}

func main() {
	tools.InitRedis()
	r := gin.Default()
	r.Use(Core())
	r.POST("/api/login", doLogin)
	r.Run(":3000")
}
