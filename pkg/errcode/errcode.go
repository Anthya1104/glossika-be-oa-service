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

	BadRequest     ErrCode = "2000"
	BadHeader      ErrCode = "2001"
	BadQuery       ErrCode = "2002"
	BadURLParam    ErrCode = "2003"
	BadRequestBody ErrCode = "2004"
)

var ErrCodeMsg = map[ErrCode]string{
	BadRequest:     "Bad request",
	BadHeader:      "Bad header",
	BadQuery:       "Bad query",
	BadURLParam:    "Bad URL parameter",
	BadRequestBody: "Bad request body",
}
