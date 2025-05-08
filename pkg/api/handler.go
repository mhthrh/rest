package api

import (
	"context"
	"github.com/gin-gonic/gin"
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

	defer func() {
		c.JSON(errors.GetHttpStatus(e, c.Request.Method), e)
	}()

	ctx := context.Background()

	if err := c.BindJSON(&u); err != nil {
		e = errors.NewErrConvertData(err)
		return
	}
	e = srv.Create(ctx, &u)
}

func getUser(c *gin.Context) {
	var (
		e   *errors.Error
		key = "userName"
		u   user.User
	)

	defer func() {
		if e.Code == errors.Success().Code {
			c.JSON(errors.GetHttpStatus(e, c.Request.Method), u)
			return
		}
		c.JSON(errors.GetHttpStatus(e, c.Request.Method), e)
	}()
	userName, ok := c.GetQuery(key)
	if !ok || userName == "" {
		e = errors.NewErrKeyNotExist(key)
		return
	}
	ctx := context.Background()
	u, e = srv.GetByUserName(ctx, userName)
}
func updateUser(c *gin.Context) {
	var (
		e *errors.Error
		u user.User
	)

	defer func() {
		c.JSON(errors.GetHttpStatus(e, c.Request.Method), e)
	}()

	ctx := context.Background()

	if err := c.BindJSON(&u); err != nil {
		e = errors.NewErrConvertData(err)
		return
	}
	e = srv.Update(ctx, &u)
}
func deleteUser(c *gin.Context) {
	var (
		e   *errors.Error
		key = "userName"
	)

	defer func() {
		c.JSON(errors.GetHttpStatus(e, c.Request.Method), e)
	}()
	userName, ok := c.GetQuery(key)
	if !ok || userName == "" {
		e = errors.NewErrKeyNotExist(key)
		return
	}
	ctx := context.Background()
	e = srv.Remove(ctx, userName)

}
