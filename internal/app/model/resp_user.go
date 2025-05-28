package model

type GetUserInfoResp struct {
	CommonRes
	Data GetUserInfoRespData `json:"data"`
}

type UserLoginResp struct {
	CommonRes
	Data UserLoginRespData `json:"data"`
}

type GetUserInfoRespData struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

type UserLoginRespData struct {
	Token string `json:"token"`
}
