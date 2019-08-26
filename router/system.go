package router

import (
	"github.com/gin-gonic/gin"
	"github.com/youshanyang/youshanyang_study_go/api/system"
)

func useDefaultRouters(geng *gin.Engine, grg *gin.RouterGroup)  {
	unknownGroup := grg.Group("/")

	useUnknownSystemRouters(unknownGroup)
}
func useUnknownSystemRouters(grg *gin.RouterGroup)  {
	grg.GET("heartbeat",system.HeartBeat)
}