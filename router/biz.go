package router


import (
	"github.com/gin-gonic/gin"
	"github.com/youshanyang/youshanyang_study_go/api/business"
)

// 定义业务路由
func useBizRouters(rg *gin.RouterGroup) {
	var test_api = rg.Group("/test_api")
	{
		test_api.POST("welcome", business.GetTestContent)
	}
}
