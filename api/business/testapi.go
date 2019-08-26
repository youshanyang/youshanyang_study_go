package business

import (
	"github.com/gin-gonic/gin"
	"github.com/youshanyang/youshanyang_study_go/model"
	"net/http"
)

func GetTestContent(ctx *gin.Context)  {
	var result = model.ResponseResult{}
	result.Code=0
	result.Content ="这是我的go：Hello,world！";
	result.Msg="返回成功"

	ctx.JSON(http.StatusOK, result)
}
