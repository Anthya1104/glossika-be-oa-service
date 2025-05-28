package database

import (
	"context"
	"net/http"

	dbModel "github.com/Anthya1104/glossika-be-oa-service/internal/app/model/db"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
)

func (db *SqlDb) GetUserInfoByUserId(ctx context.Context, userId string) (userInfo dbModel.UserInfo, wrapErr errcode.WrapErr) {
	err := db.Orm.WithContext(ctx).
		Model(&dbModel.UserInfo{}).
		Where("id = ? ", userId).
		Find(&userInfo).Error

	if err != nil {
		wrapErr = errcode.WrapErr{
			HttpStatus: http.StatusInternalServerError,
			ErrCode:    errcode.DBGetUserInfoFailed,
			RawErr:     err,
		}
	}
	log.C(ctx).Debugf("GetUserInfoByUserId userId=%v, userInfo=%v, err=%v", userId, userInfo, wrapErr.RawErr)

	return
}
