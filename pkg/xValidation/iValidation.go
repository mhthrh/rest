package xValidation

import (
	"github.com/mhthrh/common_pkg/pkg/model/user"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
)

type IValidation interface {
	Create(user *user.User) *xErrors.Error
	GetByUserName(userName string) *xErrors.Error
	Update(user *user.User) *xErrors.Error
	Remove(userName string) *xErrors.Error
}
