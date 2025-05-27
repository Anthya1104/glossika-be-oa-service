package model

type GetUserInfoResp struct {
	CommonRes
	Data GetUserInfoRespData `json:"data"`
}

type GetUserInfoRespData struct {
	UserID string `json:"user_id"`
}
