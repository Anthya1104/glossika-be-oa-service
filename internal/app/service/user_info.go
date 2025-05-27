package service

import (
	"context"

	"github.com/Anthya1104/gin-base-service/internal/app/model"
	"github.com/Anthya1104/gin-base-service/pkg/errcode"
)

func GetUserInfo(ctx context.Context, userId string) (resp model.GetUserInfoResp, wrapErr errcode.WrapErr) {

	//TODO: add DB operation logic here
	resp = model.GetUserInfoResp{
		Data: model.GetUserInfoRespData{
			UserID: userId,
		},
	}

	return
}
