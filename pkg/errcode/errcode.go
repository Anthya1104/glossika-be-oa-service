package errcode

type WrapErr struct {
	HttpStatus int
	ErrCode    ErrCode
	RawErr     error
}

type ErrCode string

const (
	DBOperationFailed ErrCode = "1000"

	DBGetUserInfoFailed ErrCode = "1001"
)
