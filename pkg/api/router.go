package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() http.Handler {

	g := gin.Default()
	g.Use(checkToken())

	userGroup := g.Group("/user")
	userGroup.Use(checkToken2())
	userGroup.Handle("POST", "/create", create)
	userGroup.Handle("GET", "/get", getUser)
	userGroup.Handle("PUT", "/update", updateUser)
	userGroup.Handle("DELETE", "/delete", deleteUser)

	return g
}
