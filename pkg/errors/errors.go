package errors

import "time"

const (
	timeFormat = "2006-01-02 15:04:05.000"
	User       = "invalidUser"
	Validation = "validation"
	Loader     = "configLoader"
	Token      = "invalidToken"
)

type Error struct {
	Code      int
	ErrorType string
	Message   string
	Detail    string
	Internal  *Error
	BaseError error
	Time      time.Time
}
