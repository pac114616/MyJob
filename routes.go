package main

import (
	"github.com/gin-gonic/gin"
	"text/controller"
)

func CollectRoute(r*gin.Engine)*gin.Engine/*代表返回类型*/{
	r.POST("/api/auth/register",controller.Register)
	r.POST("/api/auth/login",controller.Login)
	return r
}
