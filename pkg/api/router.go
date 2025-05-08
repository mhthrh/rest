package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restfullApi/pkg/errors"
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
		context.JSON(errors.GetHttpStatus(errors.NotImplemented(context.Request.Method), context.Request.Method), errors.NotImplemented(context.Request.Method))
	})
	g.NoMethod(func(context *gin.Context) {
		context.JSON(errors.GetHttpStatus(errors.NotImplemented(context.Request.Method), context.Request.Method), errors.NotImplemented(context.Request.Method))
	})
	return g
}
