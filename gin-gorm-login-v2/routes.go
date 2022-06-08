package main

import (
	"gin-gorm-login-v2/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	//注册
	r.POST("/register", controller.Register)
	//登录
	r.POST("/login", controller.Login)
	return r
}
