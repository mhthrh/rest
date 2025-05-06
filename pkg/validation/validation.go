package validation

import (
	"restfullApi/pkg/errors"
	"restfullApi/pkg/model/user"
	"restfullApi/util/validation"
)

type Validity struct {
}

func (v Validity) Create(user *user.User) *errors.Error {
	if !validation.MobilePhone(user.PhoneNumber) {
		return errors.NewErrMobilePhone(nil, nil)
	}
	if !validation.Email(user.Email) {
		return errors.NewErrEmailValidation(nil, nil)
	}
	if !validation.Name(user.FirstName) {
		return errors.NewErrName(nil, nil)
	}
	if !validation.Name(user.LastName) {
		return errors.NewErrName(nil, nil)
	}
	if !validation.Password(user.Password) {
		return errors.NewErrPasswordValidation(nil, nil)
	}
	return nil
}

func (v Validity) GetByUserName(userName string) *errors.Error {
	if !validation.Name(userName) {
		return errors.NewErrName(nil, nil)
	}
	return nil
}

func (v Validity) Update(user *user.User) *errors.Error {
	if !validation.MobilePhone(user.PhoneNumber) {
		return errors.NewErrMobilePhone(nil, nil)
	}
	if !validation.Email(user.Email) {
		return errors.NewErrEmailValidation(nil, nil)
	}
	if !validation.Name(user.FirstName) {
		return errors.NewErrName(nil, nil)
	}
	if !validation.Name(user.LastName) {
		return errors.NewErrName(nil, nil)
	}
	if !validation.Password(user.Password) {
		return errors.NewErrPasswordValidation(nil, nil)
	}
	return nil
}

func (v Validity) Remove(userName string) *errors.Error {
	if !validation.Name(userName) {
		return errors.NewErrName(nil, nil)
	}
	return nil
}
