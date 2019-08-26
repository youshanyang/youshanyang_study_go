package router

import (
	"github.com/gin-gonic/gin"
)

// UseRouters 定义路由
func UseRouters(geng *gin.Engine) {
	//自定义路由
	hostPath := "/"
	bizGroup := geng.Group(hostPath)
	useBizRouters(bizGroup)
	// 默认通用路由
	sysGroup := geng.Group(hostPath + "/sys")
	useDefaultRouters(geng, sysGroup)
}

