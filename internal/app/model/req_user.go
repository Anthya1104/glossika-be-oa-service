package model

import "regexp"

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
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[()[\]{}<>+\-*/?,.:;"'_\\|~` + "`" + `!@#$%^&=]`).MatchString(password)
	return hasUpper && hasLower && hasSpecial
}

func (req UserRegisterReq) ValidateRegisterEmail() bool {
	//validate email
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(req.Email)
}
