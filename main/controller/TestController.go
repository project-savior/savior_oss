package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"savior_oss/main/common"
)

func Test(context *gin.Context) {
	context.JSON(http.StatusOK, common.BuildData("request is ok"))
}
