package xValidation

import (
	"github.com/mhthrh/common_pkg/pkg/model/user"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"github.com/mhthrh/common_pkg/util/validation"
)

type Validity struct {
}

func New() IValidation {
	return Validity{}
}

func (v Validity) Create(user *user.User) *xErrors.Error {
	if !validation.MobilePhone(user.PhoneNumber) {
		return xErrors.NewErrMobilePhone(nil, nil)
	}
	if !validation.Email(user.Email) {
		return xErrors.NewErrEmailValidation(nil, nil)
	}
	if !validation.Name(user.FirstName) {
		return xErrors.NewErrName(nil, nil)
	}
	if !validation.Name(user.LastName) {
		return xErrors.NewErrName(nil, nil)
	}
	if !validation.Password(user.Password) {
		return xErrors.NewErrPasswordValidation(nil, nil)
	}
	return nil
}

func (v Validity) GetByUserName(userName string) *xErrors.Error {
	if !validation.Name(userName) {
		return xErrors.NewErrName(nil, nil)
	}
	return nil
}

func (v Validity) Update(user *user.User) *xErrors.Error {
	if !validation.MobilePhone(user.PhoneNumber) {
		return xErrors.NewErrMobilePhone(nil, nil)
	}
	if !validation.Email(user.Email) {
		return xErrors.NewErrEmailValidation(nil, nil)
	}
	if !validation.Name(user.FirstName) {
		return xErrors.NewErrName(nil, nil)
	}
	if !validation.Name(user.LastName) {
		return xErrors.NewErrName(nil, nil)
	}
	if !validation.Password(user.Password) {
		return xErrors.NewErrPasswordValidation(nil, nil)
	}
	return nil
}

func (v Validity) Remove(userName string) *xErrors.Error {
	if !validation.Name(userName) {
		return xErrors.NewErrName(nil, nil)
	}
	return nil
}
