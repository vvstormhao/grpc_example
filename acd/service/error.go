package service

import "errors"

var (
	ErrInvalidParam = errors.New("invalid param")
	ErrDBOperationFailed = errors.New("db operation failed")
)
