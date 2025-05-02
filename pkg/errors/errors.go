package errors

import "time"

const timeFormat = "2006-01-02 15:04:05.000"

type Error struct {
	Code      int
	Message   string
	Detail    string
	Internal  *Error
	BaseError error
	Time      time.Time
}
