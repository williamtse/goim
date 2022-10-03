package apiAuth

import (
	"GoIM/pkg/auth"
	"GoIM/pkg/common/log"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录获取token
func GetAccessToken(c *gin.Context) {
	log.NewInfo("do login")
	uid := c.PostForm("id")
	fmt.Println("form data uid:", uid)
	appId := "appid"
	appSecret := "appsecret"
	accessToken, err := auth.Login(appId, appSecret, uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"code": 0,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": accessToken,
			"code": 1,
		})
	}
}
