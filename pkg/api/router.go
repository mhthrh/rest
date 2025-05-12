package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"net/http"
)

func Run() http.Handler {
	g := gin.Default()
	g.Use(checkToken())

	userGroup := g.Group("/user")
	userGroup.Use(checkAddress())

	userGroup.POST("/create", create)
	userGroup.GET("/get", getUser)
	userGroup.PUT("/update", updateUser)
	userGroup.DELETE("/delete", deleteUser)

	g.NoRoute(func(context *gin.Context) {
		context.JSON(xErrors.GetHttpStatus(xErrors.NotImplemented(context.Request.Method), context.Request.Method), xErrors.NotImplemented(context.Request.Method))
	})
	g.NoMethod(func(context *gin.Context) {
		context.JSON(xErrors.GetHttpStatus(xErrors.NotImplemented(context.Request.Method), context.Request.Method), xErrors.NotImplemented(context.Request.Method))
	})
	return g
}
