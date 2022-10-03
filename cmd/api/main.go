package main

import (
	apiAuth "GoIM/internal/api/auth"
	apiMsg "GoIM/internal/api/msg"
	"GoIM/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(utils.CorsHandler())
	//certificate
	authRouterGroup := r.Group("/auth")
	{
		authRouterGroup.POST("/access_token", apiAuth.GetAccessToken)
	}

	imGroup := r.Group("/im")
	{
		imGroup.POST("send_msg", apiMsg.SendMsg)
	}

	r.Run(":3000")
}
