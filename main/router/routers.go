package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"savior_oss/main/common"
	"savior_oss/main/controller"
)

func SetUpRoute() *gin.Engine {
	r := gin.Default()
	registerRoute(r)
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, common.BuildResponse(404, "Not Found", nil))
	})
	return r
}

func registerRoute(r *gin.Engine) {
	r.Use()
	r.GET("/test", controller.Test)

}
