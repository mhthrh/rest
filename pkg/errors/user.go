package errors

import "time"

func NewErrUsrExist(err error, Err *Error) *Error {
	return &Error{
		Code:      100001,
		Message:   "already user exist",
		Detail:    "already user exist",
		Internal:  Err,
		BaseError: err,
		Time:      time.Now(),
	}
}
func NewErrUsrNotExist(err error, Err *Error) *Error {
	return &Error{
		Code:      100002,
		Message:   "user doesnt exist",
		Detail:    "user doesnt exist",
		Internal:  Err,
		BaseError: err,
		Time:      time.Now(),
	}
}
