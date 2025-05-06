package validation

import (
	"restfullApi/pkg/errors"
	"restfullApi/pkg/model/user"
)

type IValidation interface {
	Create(user *user.User) *errors.Error
	GetByUserName(userName string) *errors.Error
	Update(user *user.User) *errors.Error
	Remove(userName string) *errors.Error
}
