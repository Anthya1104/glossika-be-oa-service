package service

import (
	"context"
	"fmt"

	"github.com/Anthya1104/gin-base-service/internal/app/database"
	"github.com/Anthya1104/gin-base-service/internal/app/model"
	"github.com/Anthya1104/gin-base-service/pkg/errcode"
)

func GetUserInfo(ctx context.Context, userId string) (resp model.GetUserInfoResp, wrapErr errcode.WrapErr) {

	//TODO: add DB operation logic here
	userInfo, wrapErr := database.GetSqlDb().Repo.GetUserInfoByUserId(ctx, userId)
	if wrapErr.RawErr != nil {
		err := fmt.Errorf("failed to get user, userId: %v, err: %v", userId, wrapErr.RawErr)
		wrapErr.RawErr = err
		return
	}

	resp = model.GetUserInfoResp{
		Data: model.GetUserInfoRespData{
			UserID:   userInfo.Id,
			UserName: userInfo.UserName,
		},
	}

	return
}
