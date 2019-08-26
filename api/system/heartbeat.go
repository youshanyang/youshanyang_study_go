package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HeartBeat(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,&gin.H{
		"Code":"1",
		"Content":"系统运行正常！",
		"Msg":"返回成功",
	})
}