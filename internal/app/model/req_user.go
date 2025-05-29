package model

import "regexp"

// compile regexps once at package level
var (
	reUpper   = regexp.MustCompile(`[A-Z]`)
	reLower   = regexp.MustCompile(`[a-z]`)
	reSpecial = regexp.MustCompile(`[()[\]{}<>+\-*/?,.:;"'_\\|~` + "`" + `!@#$%^&=]`)
	reEmail   = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

type UserRegisterReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=60"`
}

func (req UserRegisterReq) ValidateRegisterPassword() bool {
	//validate password
	password := req.Password
	if len(password) < 6 || len(password) > 16 {
		return false
	}
	return reUpper.MatchString(password) && reLower.MatchString(password) && reSpecial.MatchString(password)
}

func (req UserRegisterReq) ValidateRegisterEmail() bool {
	//validate email
	return reEmail.MatchString(req.Email)
}

type UserLoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
