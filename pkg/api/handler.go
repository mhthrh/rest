package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"restfullApi/pkg/errors"
	"restfullApi/pkg/model/user"
	"restfullApi/pkg/service"
)

var (
	srv *service.Service
)

func init() {
	srv = service.New()
}

func create(c *gin.Context) {
	var (
		e *errors.Error
		u user.User
	)

	defer c.JSON(err2code(e), result(e))

	ctx := context.Background()

	if err := c.BindJSON(&u); err != nil {
		fmt.Println(err)
	}
	e = srv.Create(ctx, &u)

}

func getUser(c *gin.Context) {
	log.Println("getUser")

}
func updateUser(c *gin.Context) {
	log.Println("updateUser")
}
func deleteUser(c *gin.Context) {
	log.Println("deleteUser")

}
