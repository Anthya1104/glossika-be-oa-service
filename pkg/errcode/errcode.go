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

	DBCountUserFailed  ErrCode = "1011"
	DBDuplicatedUser   ErrCode = "1012"
	DBCreateUserFailed ErrCode = "1013"
	DBUpdateUserFailed ErrCode = "1014"
	DBUserNotFound     ErrCode = "1015"

	DBGetUserRecommendationFailed ErrCode = "1020"

	DBGetProductFailed ErrCode = "1030"

	BadRequest     ErrCode = "2000"
	BadHeader      ErrCode = "2001"
	BadQuery       ErrCode = "2002"
	BadURLParam    ErrCode = "2003"
	BadRequestBody ErrCode = "2004"

	UserInvalidAuth  ErrCode = "3000"
	UserNotActivated ErrCode = "3001"

	JWTGenerateFailed ErrCode = "4000"

	BcryptHashFailed ErrCode = "5000"
)

var ErrCodeMsg = map[ErrCode]string{
	BadRequest:     "Bad request",
	BadHeader:      "Bad header",
	BadQuery:       "Bad query",
	BadURLParam:    "Bad URL parameter",
	BadRequestBody: "Bad request body",
}
